// VectorAdd: out[i] = in[i] + 1.0f
//
// The checked-in kernels.hsaco is derived from the legacy ReLU kernel object
// (global_load + v_add_f32 + global_store) because NaviSim's timing model
// targets that code-object format and SMEM encoding. `make regenerate`
// reproduces the code object and embedded Go data without a GPU compiler.
__kernel void VectorAdd(
    const int count,
    __global const float* in,
    __global float* out)
{
    int index = get_global_id(0);
    if (index < count) {
        out[index] = in[index] + 1.0f;
    }
}
