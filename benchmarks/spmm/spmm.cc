// Adapted from code written by Guyue Huang for paper on GE-SpMM
// That paper is arXiv:2007.03179 [cs.DC]
#include "hip/hip_runtime.h"
#define TYPE float 

__global__ void spmm_cwm( 
        int A_nrows, int B_ncols, int* A_csrRowPtr, 
        int* A_csrColInd, TYPE* A_csrVal, TYPE* B_dnVal, 
        TYPE* C_dnVal 
) { 
    HIP_DYNAMIC_SHARED( int, sh) 
    int* colInd_sh = sh; 
    TYPE* val_sh = (TYPE*)&sh[(blockDim.y * warpSize)]; 
    int shmem_offset = (threadIdx.y * warpSize); 
    int thread_idx = shmem_offset + threadIdx.x; 

    int rid = blockDim.y * blockIdx.x + threadIdx.y; 

    if (rid < A_nrows) { 
        int cid = (blockIdx.y * warpSize * 2) + threadIdx.x; 
        int lb = A_csrRowPtr[rid]; 
        int hb = A_csrRowPtr[(rid + 1)]; 
        int ptr = lb + threadIdx.x; 
        int offset; 
        TYPE acc1 = 0, acc2 = 0, val; 

        if (blockIdx.y != gridDim.y - 1) { 
            for (int jj = lb; jj < hb; jj += warpSize) { 
                if (ptr < hb) { 
                    val_sh[thread_idx] = A_csrVal[ptr]; 
                    colInd_sh[thread_idx] = B_ncols * A_csrColInd[ptr]; 
                } 
                __syncthreads(); 
                ptr += warpSize; 

                for (int kk = 0; kk < warpSize && jj + kk < hb; kk++) { 
                    offset = colInd_sh[(shmem_offset + kk)] + cid; 
                    val = val_sh[(shmem_offset + kk)]; 
                    acc1 += val * B_dnVal[offset]; 
                    acc2 += val * B_dnVal[offset + warpSize]; 
                } 
                __syncthreads(); 
            } 
            offset = rid * B_ncols + cid; 
            C_dnVal[offset] = acc1; 
            C_dnVal[offset + warpSize] = acc2; 
        } else { 
            int nout = (B_ncols - cid + (warpSize-1)) / warpSize; 
            for (int jj = lb; jj < hb; jj += warpSize) { 
                if (ptr < hb) { 
                    val_sh[thread_idx] = A_csrVal[ptr]; 
                    colInd_sh[thread_idx] = B_ncols * A_csrColInd[ptr]; 
                } 
                __syncthreads(); 
                ptr += warpSize; 

                for (int kk = 0; kk < warpSize && jj + kk < hb; kk++) { 
                    val = val_sh[(shmem_offset + kk)]; 
                    offset = colInd_sh[(shmem_offset + kk)] + cid; 
                    if (nout > 0) { 
                        acc1 += val * B_dnVal[offset]; 
                    } if (nout > 1) { 
                        acc2 += val * B_dnVal[offset + warpSize]; 
                    } 
                } 
                __syncthreads(); 
            } 
            offset = rid * B_ncols + cid; 
            if (nout > 0) { 
                C_dnVal[offset] = acc1; 
            } 
            if (nout > 1) { 
                C_dnVal[(offset + warpSize)] = acc2; 
            } 
        } 
    } 
}
