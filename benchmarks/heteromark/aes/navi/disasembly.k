
kernels.hsaco:	file format elf64-amdgpu


Disassembly of section .text:

0000000000001000 <Encrypt$local>:
	.long 0x00000001                                           // 000000001000: 00000001
	.long 0x00000002                                           // 000000001004: 00000002
	.long 0x000a0001                                           // 000000001008: 000A0001
	.long 0x00000001                                           // 00000000100C: 00000001
	.long 0x00000100                                           // 000000001010: 00000100
		...
	.long 0x60af00c7                                           // 000000001030: 60AF00C7
	.long 0x00000090                                           // 000000001034: 00000090
	.long 0x000a040b                                           // 000000001038: 000A040B
		...
	.long 0x00000050                                           // 000000001048: 00000050
		...
	.long 0x00400019                                           // 000000001054: 00400019
		...
	.long 0x05040404                                           // 000000001064: 05040404
	.long 0xffffffff                                           // 000000001068: FFFFFFFF
		...
	s_load_dword s0, s[4:5], 0x4                               // 000000001100: F4000002 FA000004
	s_load_dword s4, s[6:7], 0x18                              // 000000001108: F4000103 FA000018
	v_mov_b32_e32 v2, 0xff                                     // 000000001110: 7E0402FF 000000FF
	s_waitcnt lgkmcnt(0)                                       // 000000001118: BF8CC07F
	s_and_b32 s5, s0, 0xffff                                   // 00000000111C: 8705FF00 0000FFFF
	s_load_dwordx4 s[0:3], s[6:7], 0x0                         // 000000001124: F4080003 FA000000
	s_mul_i32 s8, s8, s5                                       // 00000000112C: 93080508
	s_add_i32 s4, s4, s8                                       // 000000001130: 81040804
	v_add_lshl_u32 v0, s4, v0, 4                               // 000000001134: D7470000 02120004
	s_load_dwordx4 s[4:7], s[6:7], 0x10                        // 00000000113C: F4080103 FA000010
	s_waitcnt lgkmcnt(0)                                       // 000000001144: BF8CC07F
	s_mov_b64 s[6:7], 0                                        // 000000001148: BE860480
	v_ashrrev_i32_e32 v1, 31, v0                               // 00000000114C: 3002009F
	v_add_co_u32_e64 v0, vcc_lo, s0, v0                        // 000000001150: D70F6A00 00020000
	s_load_dwordx4 s[8:11], s[2:3], 0x0                        // 000000001158: F4080201 FA000000
	s_movk_i32 s0, 0xff                                        // 000000001160: B00000FF
	v_add_co_ci_u32_e32 v1, vcc_lo, s1, v1, vcc_lo             // 000000001164: 50020201
	global_load_dwordx4 v[3:6], v[0:1], off                    // 000000001168: DC388000 037D0000
	s_waitcnt lgkmcnt(0)                                       // 000000001170: BF8CC07F
	s_lshr_b32 s1, s8, 24                                      // 000000001174: 90019808
	s_lshr_b32 s13, s8, 8                                      // 000000001178: 900D8808
	s_lshr_b32 s14, s9, 24                                     // 00000000117C: 900E9809
	s_lshr_b32 s16, s9, 8                                      // 000000001180: 90108809
	s_lshr_b32 s17, s10, 24                                    // 000000001184: 9011980A
	s_lshr_b32 s19, s10, 8                                     // 000000001188: 9013880A
	s_lshr_b32 s20, s11, 24                                    // 00000000118C: 9014980B
	s_lshr_b32 s22, s11, 8                                     // 000000001190: 9016880B
	s_lshr_b32 s12, s8, 16                                     // 000000001194: 900C9008
	s_lshr_b32 s15, s9, 16                                     // 000000001198: 900F9009
	s_lshr_b32 s18, s10, 16                                    // 00000000119C: 9012900A
	s_lshr_b32 s21, s11, 16                                    // 0000000011A0: 9015900B
	s_waitcnt vmcnt(0)                                         // 0000000011A4: BF8C3F70
	v_lshrrev_b32_e32 v7, 8, v3                                // 0000000011A8: 2C0E0688
	v_lshrrev_b32_e32 v8, 8, v4                                // 0000000011AC: 2C100888
	v_lshrrev_b32_e32 v9, 8, v5                                // 0000000011B0: 2C120A88
	v_lshrrev_b32_e32 v10, 8, v6                               // 0000000011B4: 2C140C88
	v_xor_b32_e32 v34, s1, v3                                  // 0000000011B8: 3A440601
	v_xor_b32_sdwa v33, s13, v3 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 0000000011BC: 3A4206F9 0586060D
	v_xor_b32_sdwa v32, s8, v3 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 0000000011C4: 3A4006F9 03860608
	v_xor_b32_e32 v15, s14, v4                                 // 0000000011CC: 3A1E080E
	v_xor_b32_sdwa v29, s16, v4 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 0000000011D0: 3A3A08F9 05860610
	v_xor_b32_sdwa v28, s9, v4 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 0000000011D8: 3A3808F9 03860609
	v_xor_b32_e32 v27, s17, v5                                 // 0000000011E0: 3A360A11
	v_xor_b32_sdwa v25, s19, v5 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 0000000011E4: 3A320AF9 05860613
	v_xor_b32_sdwa v24, s10, v5 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 0000000011EC: 3A300AF9 0386060A
	v_xor_b32_e32 v23, s20, v6                                 // 0000000011F4: 3A2E0C14
	v_xor_b32_sdwa v21, s22, v6 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 0000000011F8: 3A2A0CF9 05860616
	v_xor_b32_sdwa v20, s11, v6 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 000000001200: 3A280CF9 0386060B
	v_xor_b32_e32 v35, s12, v7                                 // 000000001208: 3A460E0C
	v_xor_b32_e32 v31, s15, v8                                 // 00000000120C: 3A3E100F
	v_xor_b32_e32 v17, s18, v9                                 // 000000001210: 3A221212
	v_xor_b32_e32 v47, s21, v10                                // 000000001214: 3A5E1415
	v_mov_b32_e32 v3, 0                                        // 000000001218: 7E060280
	s_branch BB0_2                                             // 00000000121C: BF8200B3

0000000000001220 <BB0_1>:
	s_add_u32 s8, s2, s6                                       // 000000001220: 80080602
	s_addc_u32 s9, s3, s7                                      // 000000001224: 82090703
	s_waitcnt vmcnt(15)                                        // 000000001228: BF8C3F7F
	v_cmp_gt_i16_e32 vcc_lo, 0, v61                            // 00000000122C: 7D187A80
	s_waitcnt vmcnt(5)                                         // 000000001230: BF8C3F75
	v_lshlrev_b16_e64 v26, 1, v9                               // 000000001234: D714001A 00021281
	s_waitcnt vmcnt(0)                                         // 00000000123C: BF8C3F70
	v_lshlrev_b16_e64 v27, 1, v19                              // 000000001240: D714001B 00022681
	v_lshlrev_b16_e64 v29, 1, v7                               // 000000001248: D714001D 00020E81
	v_mov_b32_e32 v21, s9                                      // 000000001250: 7E2A0209
	v_mov_b32_e32 v20, s8                                      // 000000001254: 7E280208
	v_xor_b32_e32 v48, 27, v26                                 // 000000001258: 3A60349B
	v_xor_b32_e32 v49, 27, v27                                 // 00000000125C: 3A62369B
	v_lshlrev_b16_e64 v30, 1, v63                              // 000000001260: D714001E 00027E81
	v_xor_b32_e32 v50, 27, v29                                 // 000000001268: 3A643A9B
	global_load_dwordx4 v[22:25], v[20:21], off offset:16      // 00000000126C: DC388010 167D0014
	v_lshlrev_b16_e64 v20, 1, v61                              // 000000001274: D7140014 00027A81
	v_lshlrev_b16_e64 v21, 1, v13                              // 00000000127C: D7140015 00021A81
	v_lshlrev_b16_e64 v31, 1, v16                              // 000000001284: D714001F 00022081
	v_xor_b32_e32 v51, 27, v30                                 // 00000000128C: 3A663C9B
	v_lshlrev_b16_e64 v32, 1, v6                               // 000000001290: D7140020 00020C81
	v_xor_b32_e32 v46, 27, v20                                 // 000000001298: 3A5C289B
	v_xor_b32_e32 v47, 27, v21                                 // 00000000129C: 3A5E2A9B
	v_xor_b32_e32 v52, 27, v31                                 // 0000000012A0: 3A683E9B
	v_lshlrev_b16_e64 v34, 1, v11                              // 0000000012A4: D7140022 00021681
	v_xor_b32_e32 v53, 27, v32                                 // 0000000012AC: 3A6A409B
	v_cndmask_b32_e32 v20, v20, v46, vcc_lo                    // 0000000012B0: 02285D14
	v_cmp_gt_i16_e32 vcc_lo, 0, v13                            // 0000000012B4: 7D181A80
	v_lshlrev_b16_e64 v35, 1, v59                              // 0000000012B8: D7140023 00027681
	v_xor_b32_e32 v15, 27, v34                                 // 0000000012C0: 3A1E449B
	v_lshlrev_b16_e64 v36, 1, v57                              // 0000000012C4: D7140024 00027281
	v_lshlrev_b16_e64 v37, 1, v10                              // 0000000012CC: D7140025 00021481
	v_cndmask_b32_e32 v21, v21, v47, vcc_lo                    // 0000000012D4: 022A5F15
	v_cmp_gt_i16_e32 vcc_lo, 0, v9                             // 0000000012D8: 7D181280
	v_xor_b32_e32 v18, 27, v35                                 // 0000000012DC: 3A24469B
	v_xor_b32_e32 v46, 27, v36                                 // 0000000012E0: 3A5C489B
	v_lshlrev_b16_e64 v39, 1, v62                              // 0000000012E4: D7140027 00027C81
	v_xor_b32_e32 v47, 27, v37                                 // 0000000012EC: 3A5E4A9B
	v_cndmask_b32_e32 v26, v26, v48, vcc_lo                    // 0000000012F0: 0234611A
	v_cmp_gt_i16_e32 vcc_lo, 0, v19                            // 0000000012F4: 7D182680
	v_lshlrev_b16_e64 v40, 1, v5                               // 0000000012F8: D7140028 00020A81
	v_xor_b32_e32 v48, 27, v39                                 // 000000001300: 3A604E9B
	v_lshlrev_b16_e64 v41, 1, v56                              // 000000001304: D7140029 00027081
	v_xor_b32_e32 v33, v16, v6                                 // 00000000130C: 3A420D10
	v_cndmask_b32_e32 v27, v27, v49, vcc_lo                    // 000000001310: 0236631B
	v_cmp_gt_i16_e32 vcc_lo, 0, v7                             // 000000001314: 7D180E80
	v_xor_b32_e32 v49, 27, v40                                 // 000000001318: 3A62509B
	v_lshlrev_b16_e64 v42, 1, v58                              // 00000000131C: D714002A 00027481
	v_xor_b32_e32 v43, v58, v56                                // 000000001324: 3A56713A
	v_xor_b32_e32 v44, v13, v61                                // 000000001328: 3A587B0D
	v_cndmask_b32_e32 v29, v29, v50, vcc_lo                    // 00000000132C: 023A651D
	v_cmp_gt_i16_e32 vcc_lo, 0, v63                            // 000000001330: 7D187E80
	v_xor_b32_e32 v50, 27, v41                                 // 000000001334: 3A64529B
	v_xor_b32_e32 v45, v62, v5                                 // 000000001338: 3A5A0B3E
	v_xor_b32_e32 v38, v10, v57                                // 00000000133C: 3A4C730A
	v_xor_b32_e32 v20, v13, v20                                // 000000001340: 3A28290D
	v_cndmask_b32_e32 v30, v30, v51, vcc_lo                    // 000000001344: 023C671E
	v_cmp_gt_i16_e32 vcc_lo, 0, v16                            // 000000001348: 7D182080
	v_xor_b32_e32 v51, 27, v42                                 // 00000000134C: 3A66549B
	v_xor_b32_e32 v28, v19, v9                                 // 000000001350: 3A381313
	v_xor_b32_e32 v45, v45, v56                                // 000000001354: 3A5A712D
	s_add_u32 s6, s6, 16                                       // 000000001358: 80069006
	v_cndmask_b32_e32 v31, v31, v52, vcc_lo                    // 00000000135C: 023E691F
	v_cmp_gt_i16_e32 vcc_lo, 0, v6                             // 000000001360: 7D180C80
	s_addc_u32 s7, s7, 0                                       // 000000001364: 82078007
	s_mov_b32 s1, 0                                            // 000000001368: BE810380
	v_cndmask_b32_e32 v32, v32, v53, vcc_lo                    // 00000000136C: 02406B20
	v_cmp_gt_i16_e32 vcc_lo, 0, v11                            // 000000001370: 7D181680
	v_xor_b32_e32 v32, v32, v7                                 // 000000001374: 3A400F20
	v_cndmask_b32_e32 v17, v34, v15, vcc_lo                    // 000000001378: 02221F22
	v_cmp_gt_i16_e32 vcc_lo, 0, v59                            // 00000000137C: 7D187680
	v_cndmask_b32_e32 v53, v35, v18, vcc_lo                    // 000000001380: 026A2523
	v_cmp_gt_i16_e32 vcc_lo, 0, v57                            // 000000001384: 7D187280
	v_cndmask_b32_e32 v36, v36, v46, vcc_lo                    // 000000001388: 02485D24
	v_cmp_gt_i16_e32 vcc_lo, 0, v10                            // 00000000138C: 7D181480
	v_xor_b32_e32 v46, v33, v30                                // 000000001390: 3A5C3D21
	v_xor_b32_e32 v36, v11, v36                                // 000000001394: 3A48490B
	v_cndmask_b32_e32 v37, v37, v47, vcc_lo                    // 000000001398: 024A5F25
	v_cmp_gt_i16_e32 vcc_lo, 0, v62                            // 00000000139C: 7D187C80
	v_xor_b32_e32 v47, v26, v61                                // 0000000013A0: 3A5E7B1A
	v_xor_b32_e32 v26, v44, v26                                // 0000000013A4: 3A34352C
	v_xor_b32_e32 v44, v63, v29                                // 0000000013A8: 3A583B3F
	v_xor_b32_e32 v29, v32, v29                                // 0000000013AC: 3A3A3B20
	v_cndmask_b32_e32 v39, v39, v48, vcc_lo                    // 0000000013B0: 024E6127
	v_cmp_gt_i16_e32 vcc_lo, 0, v5                             // 0000000013B4: 7D180A80
	v_xor_b32_e32 v48, v37, v11                                // 0000000013B8: 3A601725
	v_xor_b32_e32 v30, v44, v30                                // 0000000013BC: 3A3C3D2C
	v_xor_b32_e32 v44, v46, v7                                 // 0000000013C0: 3A580F2E
	v_xor_b32_e32 v26, v26, v19                                // 0000000013C4: 3A34271A
	v_cndmask_b32_e32 v40, v40, v49, vcc_lo                    // 0000000013C8: 02506328
	v_cmp_gt_i16_e32 vcc_lo, 0, v56                            // 0000000013CC: 7D187080
	v_xor_b32_e32 v46, v48, v57                                // 0000000013D0: 3A5C7330
	v_xor_b32_e32 v32, v32, v6                                 // 0000000013D4: 3A400D20
	v_xor_b32_e32 v29, v29, v63                                // 0000000013D8: 3A3A7F1D
	v_xor_b32_e32 v40, v43, v40                                // 0000000013DC: 3A50512B
	v_cndmask_b32_e32 v41, v41, v50, vcc_lo                    // 0000000013E0: 02526529
	v_xor_b32_e32 v43, v59, v17                                // 0000000013E4: 3A56233B
	v_cmp_gt_i16_e32 vcc_lo, 0, v58                            // 0000000013E8: 7D187480
	v_xor_b32_e32 v50, v20, v21                                // 0000000013EC: 3A642B14
	v_xor_b32_e32 v21, v47, v21                                // 0000000013F0: 3A2A2B2F
	v_xor_b32_e32 v15, v62, v41                                // 0000000013F4: 3A1E533E
	v_xor_b32_e32 v20, v20, v61                                // 0000000013F8: 3A287B14
	v_cndmask_b32_e32 v42, v42, v51, vcc_lo                    // 0000000013FC: 0254672A
	v_xor_b32_e32 v43, v43, v38                                // 000000001400: 3A564D2B
	v_xor_b32_e32 v38, v36, v38                                // 000000001404: 3A4C4D24
	v_xor_b32_e32 v47, v15, v5                                 // 000000001408: 3A5E0B0F
	v_xor_b32_e32 v36, v36, v10                                // 00000000140C: 3A481524
	v_xor_b32_e32 v41, v40, v41                                // 000000001410: 3A525328
	v_xor_b32_e32 v20, v20, v9                                 // 000000001414: 3A281314
	v_xor_b32_e32 v45, v45, v42                                // 000000001418: 3A5A552D
	v_xor_b32_e32 v34, v46, v17                                // 00000000141C: 3A44232E
	v_xor_b32_e32 v38, v38, v53                                // 000000001420: 3A4C6B26
	v_xor_b32_e32 v40, v40, v5                                 // 000000001424: 3A500B28
	v_xor_b32_e32 v48, v50, v28                                // 000000001428: 3A603932
	v_xor_b32_e32 v21, v21, v28                                // 00000000142C: 3A2A3915
	v_xor_b32_e32 v28, v44, v31                                // 000000001430: 3A383F2C
	v_xor_b32_e32 v44, v26, v27                                // 000000001434: 3A58371A
	v_xor_b32_e32 v26, v30, v33                                // 000000001438: 3A34431E
	v_xor_b32_e32 v30, v32, v63                                // 00000000143C: 3A3C7F20
	v_xor_b32_e32 v32, v36, v37                                // 000000001440: 3A404B24
	v_xor_b32_e32 v35, v47, v58                                // 000000001444: 3A46752F
	v_xor_b32_e32 v55, v41, v62                                // 000000001448: 3A6E7D29
	v_xor_b32_e32 v33, v43, v53                                // 00000000144C: 3A426B2B
	v_xor_b32_e32 v36, v40, v39                                // 000000001450: 3A484F28
	v_xor_b32_e32 v51, v45, v39                                // 000000001454: 3A664F2D
	v_xor_b32_e32 v41, v35, v42                                // 000000001458: 3A525523
	v_xor_b32_e32 v20, v20, v27                                // 00000000145C: 3A283714
	v_xor_b32_e32 v39, v30, v31                                // 000000001460: 3A4E3F1E
	v_xor_b32_e32 v46, v29, v16                                // 000000001464: 3A5C211D
	v_xor_b32_e32 v40, v32, v59                                // 000000001468: 3A507720
	v_xor_b32_e32 v29, v34, v59                                // 00000000146C: 3A3A7722
	s_waitcnt vmcnt(0)                                         // 000000001470: BF8C3F70
	v_xor_b32_sdwa v35, v21, v22 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 000000001474: 3A462CF9 05060615
	v_xor_b32_sdwa v17, v38, v24 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 00000000147C: 3A2230F9 05060626
	v_lshrrev_b32_e32 v42, 8, v22                              // 000000001484: 2C542C88
	v_lshrrev_b32_e32 v21, 8, v23                              // 000000001488: 2C2A2E88
	v_lshrrev_b32_e32 v43, 8, v24                              // 00000000148C: 2C563088
	v_lshrrev_b32_e32 v38, 8, v25                              // 000000001490: 2C4C3288
	v_xor_b32_sdwa v27, v33, v24 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 000000001494: 3A3630F9 03060621
	v_xor_b32_e32 v32, v20, v22                                // 00000000149C: 3A402D14
	v_xor_b32_sdwa v31, v28, v23 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 0000000014A0: 3A3E2EF9 0506061C
	v_xor_b32_e32 v24, v29, v24                                // 0000000014A8: 3A30311D
	v_xor_b32_e32 v29, v39, v21                                // 0000000014AC: 3A3A2B27
	v_xor_b32_sdwa v34, v48, v22 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 0000000014B0: 3A442CF9 03060630
	v_xor_b32_sdwa v15, v26, v23 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 0000000014B8: 3A1E2EF9 0306061A
	v_xor_b32_e32 v28, v46, v23                                // 0000000014C0: 3A382F2E
	v_xor_b32_sdwa v23, v36, v25 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 0000000014C4: 3A2E32F9 03060624
	v_xor_b32_sdwa v47, v55, v25 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 0000000014CC: 3A5E32F9 05060637
	v_xor_b32_e32 v20, v51, v25                                // 0000000014D4: 3A283333
	v_xor_b32_e32 v33, v44, v42                                // 0000000014D8: 3A42552C
	v_xor_b32_e32 v25, v40, v43                                // 0000000014DC: 3A325728
	v_xor_b32_e32 v21, v41, v38                                // 0000000014E0: 3A2A4D29
	s_andn2_b32 vcc_lo, exec_lo, s1                            // 0000000014E4: 8A6A017E
	s_cbranch_vccz BB0_4                                       // 0000000014E8: BF860069

00000000000014ec <BB0_2>:
	v_and_b32_e32 v4, s0, v34                                  // 0000000014EC: 36084400
	s_waitcnt vmcnt(12)                                        // 0000000014F0: BF8C3F7C
	v_and_b32_e32 v6, s0, v35                                  // 0000000014F4: 360C4600
	v_and_b32_e32 v8, s0, v33                                  // 0000000014F8: 36104200
	s_waitcnt vmcnt(7)                                         // 0000000014FC: BF8C3F77
	v_and_b32_e32 v11, v2, v32                                 // 000000001500: 36164102
	s_waitcnt vmcnt(1)                                         // 000000001504: BF8C3F71
	v_and_b32_e32 v16, v2, v27                                 // 000000001508: 36203702
	v_add_co_u32_e64 v4, vcc_lo, s4, v4                        // 00000000150C: D70F6A04 00020804
	v_and_b32_e32 v25, v2, v25                                 // 000000001514: 36323302
	v_and_b32_e32 v22, v2, v47                                 // 000000001518: 362C5F02
	v_and_b32_e32 v20, v2, v20                                 // 00000000151C: 36282902
	s_cmpk_eq_i32 s6, 0xd0                                     // 000000001520: B18600D0
	v_add_co_ci_u32_e32 v5, vcc_lo, s5, v3, vcc_lo             // 000000001524: 500A0605
	v_add_co_u32_e64 v6, vcc_lo, s4, v6                        // 000000001528: D70F6A06 00020C04
	v_add_co_ci_u32_e32 v7, vcc_lo, s5, v3, vcc_lo             // 000000001530: 500E0605
	v_add_co_u32_e64 v9, vcc_lo, s4, v8                        // 000000001534: D70F6A09 00021004
	v_and_b32_e32 v8, v2, v15                                  // 00000000153C: 36101F02
	v_and_b32_e32 v15, v2, v31                                 // 000000001540: 361E3F02
	v_add_co_ci_u32_e32 v10, vcc_lo, s5, v3, vcc_lo            // 000000001544: 50140605
	v_add_co_u32_e64 v11, vcc_lo, s4, v11                      // 000000001548: D70F6A0B 00021604
	v_add_co_ci_u32_e32 v12, vcc_lo, s5, v3, vcc_lo            // 000000001550: 50180605
	v_add_co_u32_e64 v13, vcc_lo, s4, v8                       // 000000001554: D70F6A0D 00021004
	v_add_co_ci_u32_e32 v14, vcc_lo, s5, v3, vcc_lo            // 00000000155C: 501C0605
	global_load_sbyte v61, v[4:5], off                         // 000000001560: DC248000 3D7D0004
	global_load_sbyte v5, v[6:7], off                          // 000000001568: DC248000 057D0006
	global_load_sbyte v57, v[9:10], off                        // 000000001570: DC248000 397D0009
	global_load_sbyte v6, v[11:12], off                        // 000000001578: DC248000 067D000B
	global_load_sbyte v7, v[13:14], off                        // 000000001580: DC248000 077D000D
	v_add_co_u32_e64 v9, vcc_lo, s4, v15                       // 000000001588: D70F6A09 00021E04
	v_and_b32_e32 v11, v2, v29                                 // 000000001590: 36163B02
	v_and_b32_e32 v13, v2, v28                                 // 000000001594: 361A3902
	v_add_co_ci_u32_e32 v10, vcc_lo, s5, v3, vcc_lo            // 000000001598: 50140605
	v_add_co_u32_e64 v11, vcc_lo, s4, v11                      // 00000000159C: D70F6A0B 00021604
	v_add_co_ci_u32_e32 v12, vcc_lo, s5, v3, vcc_lo            // 0000000015A4: 50180605
	v_add_co_u32_e64 v14, vcc_lo, s4, v13                      // 0000000015A8: D70F6A0E 00021A04
	v_and_b32_e32 v13, v2, v17                                 // 0000000015B0: 361A2302
	v_add_co_ci_u32_e32 v15, vcc_lo, s5, v3, vcc_lo            // 0000000015B4: 501E0605
	v_add_co_u32_e64 v16, vcc_lo, s4, v16                      // 0000000015B8: D70F6A10 00022004
	v_add_co_ci_u32_e32 v17, vcc_lo, s5, v3, vcc_lo            // 0000000015C0: 50220605
	v_add_co_u32_e64 v18, vcc_lo, s4, v13                      // 0000000015C4: D70F6A12 00021A04
	s_waitcnt vmcnt(5)                                         // 0000000015CC: BF8C3F75
	v_add_co_ci_u32_e32 v19, vcc_lo, s5, v3, vcc_lo            // 0000000015D0: 50260605
	global_load_sbyte v13, v[9:10], off                        // 0000000015D4: DC248000 0D7D0009
	global_load_sbyte v56, v[11:12], off                       // 0000000015DC: DC248000 387D000B
	global_load_sbyte v10, v[14:15], off                       // 0000000015E4: DC248000 0A7D000E
	global_load_sbyte v11, v[16:17], off                       // 0000000015EC: DC248000 0B7D0010
	global_load_sbyte v63, v[18:19], off                       // 0000000015F4: DC248000 3F7D0012
	v_add_co_u32_e64 v14, vcc_lo, s4, v25                      // 0000000015FC: D70F6A0E 00023204
	v_and_b32_e32 v16, v2, v24                                 // 000000001604: 36203102
	v_and_b32_e32 v18, v2, v23                                 // 000000001608: 36242F02
	v_and_b32_e32 v23, v2, v21                                 // 00000000160C: 362E2B02
	v_add_co_ci_u32_e32 v15, vcc_lo, s5, v3, vcc_lo            // 000000001610: 501E0605
	v_add_co_u32_e64 v16, vcc_lo, s4, v16                      // 000000001614: D70F6A10 00022004
	v_add_co_ci_u32_e32 v17, vcc_lo, s5, v3, vcc_lo            // 00000000161C: 50220605
	v_add_co_u32_e64 v18, vcc_lo, s4, v18                      // 000000001620: D70F6A12 00022404
	v_add_co_ci_u32_e32 v19, vcc_lo, s5, v3, vcc_lo            // 000000001628: 50260605
	v_add_co_u32_e64 v21, vcc_lo, s4, v22                      // 00000000162C: D70F6A15 00022C04
	v_add_co_ci_u32_e32 v22, vcc_lo, s5, v3, vcc_lo            // 000000001634: 502C0605
	v_add_co_u32_e64 v23, vcc_lo, s4, v23                      // 000000001638: D70F6A17 00022E04
	v_add_co_ci_u32_e32 v24, vcc_lo, s5, v3, vcc_lo            // 000000001640: 50300605
	v_add_co_u32_e64 v25, vcc_lo, s4, v20                      // 000000001644: D70F6A19 00022804
	v_add_co_ci_u32_e32 v26, vcc_lo, s5, v3, vcc_lo            // 00000000164C: 50340605
	global_load_sbyte v9, v[14:15], off                        // 000000001650: DC248000 097D000E
	global_load_sbyte v58, v[16:17], off                       // 000000001658: DC248000 3A7D0010
	global_load_sbyte v62, v[18:19], off                       // 000000001660: DC248000 3E7D0012
	global_load_sbyte v59, v[21:22], off                       // 000000001668: DC248000 3B7D0015
	global_load_sbyte v16, v[23:24], off                       // 000000001670: DC248000 107D0017
	global_load_sbyte v19, v[25:26], off                       // 000000001678: DC248000 137D0019
	s_cbranch_scc0 BB0_1                                       // 000000001680: BF84FEE7
	s_mov_b32 s1, -1                                           // 000000001684: BE8103C1
	s_andn2_b32 vcc_lo, exec_lo, s1                            // 000000001688: 8A6A017E
	s_cbranch_vccnz BB0_2                                      // 00000000168C: BF87FF97

0000000000001690 <BB0_4>:
	v_mov_b32_e32 v2, s2                                       // 000000001690: 7E040202
	v_mov_b32_e32 v3, s3                                       // 000000001694: 7E060203
	global_load_dwordx4 v[20:23], v[2:3], off offset:224       // 000000001698: DC3880E0 147D0002
	s_waitcnt vmcnt(0)                                         // 0000000016A0: BF8C3F70
	v_lshrrev_b32_e32 v2, 8, v20                               // 0000000016A4: 2C042888
	v_xor_b32_sdwa v8, v13, v20 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 0000000016A8: 3A1028F9 0506010D
	v_xor_b32_sdwa v3, v61, v20 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 0000000016B0: 3A0628F9 0306063D
	v_xor_b32_sdwa v13, v19, v20 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:DWORD// 0000000016B8: 3A1A28F9 06060113
	v_lshrrev_b32_e32 v19, 8, v21                              // 0000000016C0: 2C262A88
	v_lshrrev_b32_e32 v20, 8, v22                              // 0000000016C4: 2C282C88
	v_xor_b32_sdwa v7, v7, v21 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 0000000016C8: 3A0E2AF9 03060607
	v_xor_b32_sdwa v12, v63, v21 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 0000000016D0: 3A182AF9 0506013F
	v_xor_b32_sdwa v6, v6, v21 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:DWORD// 0000000016D8: 3A0C2AF9 06060106
	v_lshrrev_b32_e32 v21, 8, v23                              // 0000000016E0: 2C2A2E88
	v_xor_b32_e32 v2, v9, v2                                   // 0000000016E4: 3A040509
	v_xor_b32_e32 v19, v16, v19                                // 0000000016E8: 3A262710
	v_xor_b32_sdwa v11, v11, v22 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 0000000016EC: 3A162CF9 0306060B
	v_xor_b32_sdwa v18, v59, v22 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 0000000016F4: 3A242CF9 0506013B
	v_xor_b32_sdwa v10, v10, v22 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:DWORD// 0000000016FC: 3A142CF9 0606010A
	v_xor_b32_e32 v4, v57, v20                                 // 000000001704: 3A082939
	v_xor_b32_sdwa v27, v62, v23 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3// 000000001708: 3A362EF9 0306063E
	v_xor_b32_sdwa v5, v5, v23 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1// 000000001710: 3A0A2EF9 05060105
	v_xor_b32_sdwa v15, v58, v23 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:DWORD// 000000001718: 3A1E2EF9 0606013A
	v_xor_b32_e32 v9, v56, v21                                 // 000000001720: 3A122B38
	v_or_b32_sdwa v8, v3, v8 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD// 000000001724: 381010F9 06000603
	v_or_b32_sdwa v2, v2, v13 dst_sel:WORD_1 dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD// 00000000172C: 38041AF9 06000502
	v_or_b32_sdwa v5, v27, v5 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD// 000000001734: 380A0AF9 0600061B
	v_or_b32_sdwa v11, v11, v18 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD// 00000000173C: 381624F9 0600060B
	v_or_b32_sdwa v4, v4, v10 dst_sel:WORD_1 dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD// 000000001744: 380814F9 06000504
	v_or_b32_sdwa v3, v9, v15 dst_sel:WORD_1 dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD// 00000000174C: 38061EF9 06000509
	v_or_b32_sdwa v7, v7, v12 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD// 000000001754: 380E18F9 06000607
	v_or_b32_sdwa v6, v19, v6 dst_sel:WORD_1 dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD// 00000000175C: 380C0CF9 06000513
	v_or_b32_sdwa v2, v8, v2 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:WORD_0 src1_sel:DWORD// 000000001764: 380404F9 06040608
	v_or_b32_sdwa v4, v11, v4 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:WORD_0 src1_sel:DWORD// 00000000176C: 380808F9 0604060B
	v_or_b32_sdwa v5, v5, v3 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:WORD_0 src1_sel:DWORD// 000000001774: 380A06F9 06040605
	v_or_b32_sdwa v3, v7, v6 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:WORD_0 src1_sel:DWORD// 00000000177C: 38060CF9 06040607
	global_store_dwordx4 v[0:1], v[2:5], off                   // 000000001784: DC788000 007D0200
	s_endpgm                                                   // 00000000178C: BF810000
	s_code_end                                                 // 000000001790: BF9F0000
	s_code_end                                                 // 000000001794: BF9F0000
	s_code_end                                                 // 000000001798: BF9F0000
	s_code_end                                                 // 00000000179C: BF9F0000
	s_code_end                                                 // 0000000017A0: BF9F0000
	s_code_end                                                 // 0000000017A4: BF9F0000
	s_code_end                                                 // 0000000017A8: BF9F0000
	s_code_end                                                 // 0000000017AC: BF9F0000
	s_code_end                                                 // 0000000017B0: BF9F0000
	s_code_end                                                 // 0000000017B4: BF9F0000
	s_code_end                                                 // 0000000017B8: BF9F0000
	s_code_end                                                 // 0000000017BC: BF9F0000
	s_code_end                                                 // 0000000017C0: BF9F0000
	s_code_end                                                 // 0000000017C4: BF9F0000
	s_code_end                                                 // 0000000017C8: BF9F0000
	s_code_end                                                 // 0000000017CC: BF9F0000
	s_code_end                                                 // 0000000017D0: BF9F0000
	s_code_end                                                 // 0000000017D4: BF9F0000
	s_code_end                                                 // 0000000017D8: BF9F0000
	s_code_end                                                 // 0000000017DC: BF9F0000
	s_code_end                                                 // 0000000017E0: BF9F0000
	s_code_end                                                 // 0000000017E4: BF9F0000
	s_code_end                                                 // 0000000017E8: BF9F0000
	s_code_end                                                 // 0000000017EC: BF9F0000
	s_code_end                                                 // 0000000017F0: BF9F0000
	s_code_end                                                 // 0000000017F4: BF9F0000
	s_code_end                                                 // 0000000017F8: BF9F0000
	s_code_end                                                 // 0000000017FC: BF9F0000
	s_code_end                                                 // 000000001800: BF9F0000
	s_code_end                                                 // 000000001804: BF9F0000
	s_code_end                                                 // 000000001808: BF9F0000
	s_code_end                                                 // 00000000180C: BF9F0000
	s_code_end                                                 // 000000001810: BF9F0000
	s_code_end                                                 // 000000001814: BF9F0000
	s_code_end                                                 // 000000001818: BF9F0000
	s_code_end                                                 // 00000000181C: BF9F0000
	s_code_end                                                 // 000000001820: BF9F0000
	s_code_end                                                 // 000000001824: BF9F0000
	s_code_end                                                 // 000000001828: BF9F0000
	s_code_end                                                 // 00000000182C: BF9F0000
	s_code_end                                                 // 000000001830: BF9F0000
	s_code_end                                                 // 000000001834: BF9F0000
	s_code_end                                                 // 000000001838: BF9F0000
	s_code_end                                                 // 00000000183C: BF9F0000
	s_code_end                                                 // 000000001840: BF9F0000
	s_code_end                                                 // 000000001844: BF9F0000
	s_code_end                                                 // 000000001848: BF9F0000
	s_code_end                                                 // 00000000184C: BF9F0000
	s_code_end                                                 // 000000001850: BF9F0000
	s_code_end                                                 // 000000001854: BF9F0000
	s_code_end                                                 // 000000001858: BF9F0000
	s_code_end                                                 // 00000000185C: BF9F0000
	s_code_end                                                 // 000000001860: BF9F0000
	s_code_end                                                 // 000000001864: BF9F0000
	s_code_end                                                 // 000000001868: BF9F0000
	s_code_end                                                 // 00000000186C: BF9F0000
	s_code_end                                                 // 000000001870: BF9F0000
	s_code_end                                                 // 000000001874: BF9F0000
	s_code_end                                                 // 000000001878: BF9F0000
	s_code_end                                                 // 00000000187C: BF9F0000
