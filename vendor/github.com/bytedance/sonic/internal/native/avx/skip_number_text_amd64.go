// +build amd64
// Code generated by asm2asm, DO NOT EDIT.

package avx

var _text_skip_number = []byte{
	// .p2align 4, 0x00
	// LCPI0_0
	0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, 0x2f, // QUAD $0x2f2f2f2f2f2f2f2f; QUAD $0x2f2f2f2f2f2f2f2f  // .space 16, '////////////////'
	//0x00000010 LCPI0_1
	0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, 0x3a, //0x00000010 QUAD $0x3a3a3a3a3a3a3a3a; QUAD $0x3a3a3a3a3a3a3a3a  // .space 16, '::::::::::::::::'
	//0x00000020 LCPI0_2
	0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, 0x2b, //0x00000020 QUAD $0x2b2b2b2b2b2b2b2b; QUAD $0x2b2b2b2b2b2b2b2b  // .space 16, '++++++++++++++++'
	//0x00000030 LCPI0_3
	0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d, //0x00000030 QUAD $0x2d2d2d2d2d2d2d2d; QUAD $0x2d2d2d2d2d2d2d2d  // .space 16, '----------------'
	//0x00000040 LCPI0_4
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, //0x00000040 QUAD $0x2020202020202020; QUAD $0x2020202020202020  // .space 16, '                '
	//0x00000050 LCPI0_5
	0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, 0x2e, //0x00000050 QUAD $0x2e2e2e2e2e2e2e2e; QUAD $0x2e2e2e2e2e2e2e2e  // .space 16, '................'
	//0x00000060 LCPI0_6
	0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, 0x65, //0x00000060 QUAD $0x6565656565656565; QUAD $0x6565656565656565  // .space 16, 'eeeeeeeeeeeeeeee'
	//0x00000070 .p2align 4, 0x90
	//0x00000070 _skip_number
	0x55, //0x00000070 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000071 movq         %rsp, %rbp
	0x41, 0x57, //0x00000074 pushq        %r15
	0x41, 0x56, //0x00000076 pushq        %r14
	0x41, 0x55, //0x00000078 pushq        %r13
	0x41, 0x54, //0x0000007a pushq        %r12
	0x53, //0x0000007c pushq        %rbx
	0x48, 0x83, 0xec, 0x18, //0x0000007d subq         $24, %rsp
	0x48, 0x8b, 0x1f, //0x00000081 movq         (%rdi), %rbx
	0x4c, 0x8b, 0x6f, 0x08, //0x00000084 movq         $8(%rdi), %r13
	0x48, 0x8b, 0x16, //0x00000088 movq         (%rsi), %rdx
	0x49, 0x29, 0xd5, //0x0000008b subq         %rdx, %r13
	0x31, 0xc0, //0x0000008e xorl         %eax, %eax
	0x80, 0x3c, 0x13, 0x2d, //0x00000090 cmpb         $45, (%rbx,%rdx)
	0x4c, 0x8d, 0x3c, 0x13, //0x00000094 leaq         (%rbx,%rdx), %r15
	0x0f, 0x94, 0xc0, //0x00000098 sete         %al
	0x49, 0x01, 0xc7, //0x0000009b addq         %rax, %r15
	0x49, 0x29, 0xc5, //0x0000009e subq         %rax, %r13
	0x0f, 0x84, 0xee, 0x03, 0x00, 0x00, //0x000000a1 je           LBB0_1
	0x41, 0x8a, 0x3f, //0x000000a7 movb         (%r15), %dil
	0x8d, 0x4f, 0xd0, //0x000000aa leal         $-48(%rdi), %ecx
	0x48, 0xc7, 0xc0, 0xfe, 0xff, 0xff, 0xff, //0x000000ad movq         $-2, %rax
	0x80, 0xf9, 0x09, //0x000000b4 cmpb         $9, %cl
	0x0f, 0x87, 0xae, 0x03, 0x00, 0x00, //0x000000b7 ja           LBB0_59
	0x40, 0x80, 0xff, 0x30, //0x000000bd cmpb         $48, %dil
	0x0f, 0x85, 0x34, 0x00, 0x00, 0x00, //0x000000c1 jne          LBB0_7
	0xbf, 0x01, 0x00, 0x00, 0x00, //0x000000c7 movl         $1, %edi
	0x49, 0x83, 0xfd, 0x01, //0x000000cc cmpq         $1, %r13
	0x0f, 0x84, 0x6a, 0x03, 0x00, 0x00, //0x000000d0 je           LBB0_58
	0x41, 0x8a, 0x47, 0x01, //0x000000d6 movb         $1(%r15), %al
	0x04, 0xd2, //0x000000da addb         $-46, %al
	0x3c, 0x37, //0x000000dc cmpb         $55, %al
	0x0f, 0x87, 0x5c, 0x03, 0x00, 0x00, //0x000000de ja           LBB0_58
	0x0f, 0xb6, 0xc0, //0x000000e4 movzbl       %al, %eax
	0x48, 0xb9, 0x01, 0x00, 0x80, 0x00, 0x00, 0x00, 0x80, 0x00, //0x000000e7 movabsq      $36028797027352577, %rcx
	0x48, 0x0f, 0xa3, 0xc1, //0x000000f1 btq          %rax, %rcx
	0x0f, 0x83, 0x45, 0x03, 0x00, 0x00, //0x000000f5 jae          LBB0_58
	//0x000000fb LBB0_7
	0x48, 0x89, 0x55, 0xd0, //0x000000fb movq         %rdx, $-48(%rbp)
	0x49, 0x83, 0xfd, 0x10, //0x000000ff cmpq         $16, %r13
	0x0f, 0x82, 0x98, 0x03, 0x00, 0x00, //0x00000103 jb           LBB0_8
	0x48, 0x89, 0x5d, 0xc8, //0x00000109 movq         %rbx, $-56(%rbp)
	0x48, 0x89, 0x75, 0xc0, //0x0000010d movq         %rsi, $-64(%rbp)
	0x4d, 0x8d, 0x45, 0xf0, //0x00000111 leaq         $-16(%r13), %r8
	0x4c, 0x89, 0xc0, //0x00000115 movq         %r8, %rax
	0x48, 0x83, 0xe0, 0xf0, //0x00000118 andq         $-16, %rax
	0x4e, 0x8d, 0x64, 0x38, 0x10, //0x0000011c leaq         $16(%rax,%r15), %r12
	0x41, 0x83, 0xe0, 0x0f, //0x00000121 andl         $15, %r8d
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x00000125 movq         $-1, %r11
	0xc5, 0x7a, 0x6f, 0x05, 0xcc, 0xfe, 0xff, 0xff, //0x0000012c vmovdqu      $-308(%rip), %xmm8  /* LCPI0_0+0(%rip) */
	0xc5, 0x7a, 0x6f, 0x0d, 0xd4, 0xfe, 0xff, 0xff, //0x00000134 vmovdqu      $-300(%rip), %xmm9  /* LCPI0_1+0(%rip) */
	0xc5, 0x7a, 0x6f, 0x15, 0xdc, 0xfe, 0xff, 0xff, //0x0000013c vmovdqu      $-292(%rip), %xmm10  /* LCPI0_2+0(%rip) */
	0xc5, 0x7a, 0x6f, 0x1d, 0xe4, 0xfe, 0xff, 0xff, //0x00000144 vmovdqu      $-284(%rip), %xmm11  /* LCPI0_3+0(%rip) */
	0xc5, 0xfa, 0x6f, 0x25, 0xec, 0xfe, 0xff, 0xff, //0x0000014c vmovdqu      $-276(%rip), %xmm4  /* LCPI0_4+0(%rip) */
	0xc5, 0xfa, 0x6f, 0x2d, 0xf4, 0xfe, 0xff, 0xff, //0x00000154 vmovdqu      $-268(%rip), %xmm5  /* LCPI0_5+0(%rip) */
	0xc5, 0xfa, 0x6f, 0x35, 0xfc, 0xfe, 0xff, 0xff, //0x0000015c vmovdqu      $-260(%rip), %xmm6  /* LCPI0_6+0(%rip) */
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x00000164 movq         $-1, %r14
	0x49, 0xc7, 0xc2, 0xff, 0xff, 0xff, 0xff, //0x0000016b movq         $-1, %r10
	0x4c, 0x89, 0xfb, //0x00000172 movq         %r15, %rbx
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000175 .p2align 4, 0x90
	//0x00000180 LBB0_10
	0xc5, 0xfa, 0x6f, 0x3b, //0x00000180 vmovdqu      (%rbx), %xmm7
	0xc4, 0xc1, 0x41, 0x64, 0xc0, //0x00000184 vpcmpgtb     %xmm8, %xmm7, %xmm0
	0xc5, 0xb1, 0x64, 0xcf, //0x00000189 vpcmpgtb     %xmm7, %xmm9, %xmm1
	0xc5, 0xf9, 0xdb, 0xc1, //0x0000018d vpand        %xmm1, %xmm0, %xmm0
	0xc5, 0xa9, 0x74, 0xcf, //0x00000191 vpcmpeqb     %xmm7, %xmm10, %xmm1
	0xc5, 0xa1, 0x74, 0xd7, //0x00000195 vpcmpeqb     %xmm7, %xmm11, %xmm2
	0xc5, 0xe9, 0xeb, 0xc9, //0x00000199 vpor         %xmm1, %xmm2, %xmm1
	0xc5, 0xc1, 0xeb, 0xd4, //0x0000019d vpor         %xmm4, %xmm7, %xmm2
	0xc5, 0xe9, 0x74, 0xd6, //0x000001a1 vpcmpeqb     %xmm6, %xmm2, %xmm2
	0xc5, 0xc1, 0x74, 0xfd, //0x000001a5 vpcmpeqb     %xmm5, %xmm7, %xmm7
	0xc5, 0xe9, 0xeb, 0xdf, //0x000001a9 vpor         %xmm7, %xmm2, %xmm3
	0xc5, 0xf1, 0xeb, 0xc0, //0x000001ad vpor         %xmm0, %xmm1, %xmm0
	0xc5, 0xe1, 0xeb, 0xc0, //0x000001b1 vpor         %xmm0, %xmm3, %xmm0
	0xc5, 0xf9, 0xd7, 0xff, //0x000001b5 vpmovmskb    %xmm7, %edi
	0xc5, 0xf9, 0xd7, 0xf2, //0x000001b9 vpmovmskb    %xmm2, %esi
	0xc5, 0xf9, 0xd7, 0xc1, //0x000001bd vpmovmskb    %xmm1, %eax
	0xc5, 0xf9, 0xd7, 0xc8, //0x000001c1 vpmovmskb    %xmm0, %ecx
	0xba, 0xff, 0xff, 0xff, 0xff, //0x000001c5 movl         $4294967295, %edx
	0x48, 0x31, 0xd1, //0x000001ca xorq         %rdx, %rcx
	0x48, 0x0f, 0xbc, 0xc9, //0x000001cd bsfq         %rcx, %rcx
	0x83, 0xf9, 0x10, //0x000001d1 cmpl         $16, %ecx
	0x0f, 0x84, 0x11, 0x00, 0x00, 0x00, //0x000001d4 je           LBB0_12
	0xba, 0xff, 0xff, 0xff, 0xff, //0x000001da movl         $-1, %edx
	0xd3, 0xe2, //0x000001df shll         %cl, %edx
	0xf7, 0xd2, //0x000001e1 notl         %edx
	0x21, 0xd7, //0x000001e3 andl         %edx, %edi
	0x21, 0xd6, //0x000001e5 andl         %edx, %esi
	0x21, 0xc2, //0x000001e7 andl         %eax, %edx
	0x89, 0xd0, //0x000001e9 movl         %edx, %eax
	//0x000001eb LBB0_12
	0x44, 0x8d, 0x4f, 0xff, //0x000001eb leal         $-1(%rdi), %r9d
	0x41, 0x21, 0xf9, //0x000001ef andl         %edi, %r9d
	0x0f, 0x85, 0x26, 0x02, 0x00, 0x00, //0x000001f2 jne          LBB0_13
	0x8d, 0x56, 0xff, //0x000001f8 leal         $-1(%rsi), %edx
	0x21, 0xf2, //0x000001fb andl         %esi, %edx
	0x0f, 0x85, 0x10, 0x02, 0x00, 0x00, //0x000001fd jne          LBB0_17
	0x8d, 0x50, 0xff, //0x00000203 leal         $-1(%rax), %edx
	0x21, 0xc2, //0x00000206 andl         %eax, %edx
	0x0f, 0x85, 0x05, 0x02, 0x00, 0x00, //0x00000208 jne          LBB0_17
	0x85, 0xff, //0x0000020e testl        %edi, %edi
	0x0f, 0x84, 0x19, 0x00, 0x00, 0x00, //0x00000210 je           LBB0_22
	0x48, 0x89, 0xda, //0x00000216 movq         %rbx, %rdx
	0x4c, 0x29, 0xfa, //0x00000219 subq         %r15, %rdx
	0x0f, 0xbc, 0xff, //0x0000021c bsfl         %edi, %edi
	0x48, 0x01, 0xd7, //0x0000021f addq         %rdx, %rdi
	0x49, 0x83, 0xfa, 0xff, //0x00000222 cmpq         $-1, %r10
	0x0f, 0x85, 0xfc, 0x01, 0x00, 0x00, //0x00000226 jne          LBB0_15
	0x49, 0x89, 0xfa, //0x0000022c movq         %rdi, %r10
	//0x0000022f LBB0_22
	0x85, 0xf6, //0x0000022f testl        %esi, %esi
	0x0f, 0x84, 0x19, 0x00, 0x00, 0x00, //0x00000231 je           LBB0_25
	0x48, 0x89, 0xda, //0x00000237 movq         %rbx, %rdx
	0x4c, 0x29, 0xfa, //0x0000023a subq         %r15, %rdx
	0x0f, 0xbc, 0xfe, //0x0000023d bsfl         %esi, %edi
	0x48, 0x01, 0xd7, //0x00000240 addq         %rdx, %rdi
	0x49, 0x83, 0xfe, 0xff, //0x00000243 cmpq         $-1, %r14
	0x0f, 0x85, 0xdb, 0x01, 0x00, 0x00, //0x00000247 jne          LBB0_15
	0x49, 0x89, 0xfe, //0x0000024d movq         %rdi, %r14
	//0x00000250 LBB0_25
	0x85, 0xc0, //0x00000250 testl        %eax, %eax
	0x0f, 0x84, 0x19, 0x00, 0x00, 0x00, //0x00000252 je           LBB0_28
	0x48, 0x89, 0xda, //0x00000258 movq         %rbx, %rdx
	0x4c, 0x29, 0xfa, //0x0000025b subq         %r15, %rdx
	0x0f, 0xbc, 0xf8, //0x0000025e bsfl         %eax, %edi
	0x48, 0x01, 0xd7, //0x00000261 addq         %rdx, %rdi
	0x49, 0x83, 0xfb, 0xff, //0x00000264 cmpq         $-1, %r11
	0x0f, 0x85, 0xba, 0x01, 0x00, 0x00, //0x00000268 jne          LBB0_15
	0x49, 0x89, 0xfb, //0x0000026e movq         %rdi, %r11
	//0x00000271 LBB0_28
	0x83, 0xf9, 0x10, //0x00000271 cmpl         $16, %ecx
	0x0f, 0x85, 0xb2, 0x00, 0x00, 0x00, //0x00000274 jne          LBB0_60
	0x48, 0x83, 0xc3, 0x10, //0x0000027a addq         $16, %rbx
	0x49, 0x83, 0xc5, 0xf0, //0x0000027e addq         $-16, %r13
	0x49, 0x83, 0xfd, 0x0f, //0x00000282 cmpq         $15, %r13
	0x0f, 0x87, 0xf4, 0xfe, 0xff, 0xff, //0x00000286 ja           LBB0_10
	0x4d, 0x85, 0xc0, //0x0000028c testq        %r8, %r8
	0x48, 0x8b, 0x75, 0xc0, //0x0000028f movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x5d, 0xc8, //0x00000293 movq         $-56(%rbp), %rbx
	0x0f, 0x84, 0x9d, 0x00, 0x00, 0x00, //0x00000297 je           LBB0_42
	//0x0000029d LBB0_31
	0x4b, 0x8d, 0x04, 0x04, //0x0000029d leaq         (%r12,%r8), %rax
	0x48, 0x8d, 0x0d, 0x1c, 0x02, 0x00, 0x00, //0x000002a1 leaq         $540(%rip), %rcx  /* LJTI0_0+0(%rip) */
	0xe9, 0x0f, 0x00, 0x00, 0x00, //0x000002a8 jmp          LBB0_32
	0x90, 0x90, 0x90, //0x000002ad .p2align 4, 0x90
	//0x000002b0 LBB0_40
	0x49, 0x89, 0xd4, //0x000002b0 movq         %rdx, %r12
	0x49, 0xff, 0xc8, //0x000002b3 decq         %r8
	0x0f, 0x84, 0x8f, 0x01, 0x00, 0x00, //0x000002b6 je           LBB0_41
	//0x000002bc LBB0_32
	0x41, 0x0f, 0xbe, 0x3c, 0x24, //0x000002bc movsbl       (%r12), %edi
	0x83, 0xc7, 0xd5, //0x000002c1 addl         $-43, %edi
	0x83, 0xff, 0x3a, //0x000002c4 cmpl         $58, %edi
	0x0f, 0x87, 0x6d, 0x00, 0x00, 0x00, //0x000002c7 ja           LBB0_42
	0x49, 0x8d, 0x54, 0x24, 0x01, //0x000002cd leaq         $1(%r12), %rdx
	0x48, 0x63, 0x3c, 0xb9, //0x000002d2 movslq       (%rcx,%rdi,4), %rdi
	0x48, 0x01, 0xcf, //0x000002d6 addq         %rcx, %rdi
	0xff, 0xe7, //0x000002d9 jmpq         *%rdi
	//0x000002db LBB0_38
	0x48, 0x89, 0xd7, //0x000002db movq         %rdx, %rdi
	0x4c, 0x29, 0xff, //0x000002de subq         %r15, %rdi
	0x49, 0x83, 0xfb, 0xff, //0x000002e1 cmpq         $-1, %r11
	0x0f, 0x85, 0x95, 0x01, 0x00, 0x00, //0x000002e5 jne          LBB0_61
	0x48, 0xff, 0xcf, //0x000002eb decq         %rdi
	0x49, 0x89, 0xfb, //0x000002ee movq         %rdi, %r11
	0xe9, 0xba, 0xff, 0xff, 0xff, //0x000002f1 jmp          LBB0_40
	//0x000002f6 LBB0_36
	0x48, 0x89, 0xd7, //0x000002f6 movq         %rdx, %rdi
	0x4c, 0x29, 0xff, //0x000002f9 subq         %r15, %rdi
	0x49, 0x83, 0xfe, 0xff, //0x000002fc cmpq         $-1, %r14
	0x0f, 0x85, 0x7a, 0x01, 0x00, 0x00, //0x00000300 jne          LBB0_61
	0x48, 0xff, 0xcf, //0x00000306 decq         %rdi
	0x49, 0x89, 0xfe, //0x00000309 movq         %rdi, %r14
	0xe9, 0x9f, 0xff, 0xff, 0xff, //0x0000030c jmp          LBB0_40
	//0x00000311 LBB0_34
	0x48, 0x89, 0xd7, //0x00000311 movq         %rdx, %rdi
	0x4c, 0x29, 0xff, //0x00000314 subq         %r15, %rdi
	0x49, 0x83, 0xfa, 0xff, //0x00000317 cmpq         $-1, %r10
	0x0f, 0x85, 0x5f, 0x01, 0x00, 0x00, //0x0000031b jne          LBB0_61
	0x48, 0xff, 0xcf, //0x00000321 decq         %rdi
	0x49, 0x89, 0xfa, //0x00000324 movq         %rdi, %r10
	0xe9, 0x84, 0xff, 0xff, 0xff, //0x00000327 jmp          LBB0_40
	//0x0000032c LBB0_60
	0x48, 0x01, 0xcb, //0x0000032c addq         %rcx, %rbx
	0x49, 0x89, 0xdc, //0x0000032f movq         %rbx, %r12
	0x48, 0x8b, 0x75, 0xc0, //0x00000332 movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x5d, 0xc8, //0x00000336 movq         $-56(%rbp), %rbx
	//0x0000033a LBB0_42
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x0000033a movq         $-1, %rdi
	0x4d, 0x85, 0xf6, //0x00000341 testq        %r14, %r14
	0x0f, 0x84, 0x14, 0x01, 0x00, 0x00, //0x00000344 je           LBB0_57
	//0x0000034a LBB0_43
	0x4d, 0x85, 0xdb, //0x0000034a testq        %r11, %r11
	0x0f, 0x84, 0x0b, 0x01, 0x00, 0x00, //0x0000034d je           LBB0_57
	0x4d, 0x85, 0xd2, //0x00000353 testq        %r10, %r10
	0x48, 0x8b, 0x55, 0xd0, //0x00000356 movq         $-48(%rbp), %rdx
	0x0f, 0x84, 0xfe, 0x00, 0x00, 0x00, //0x0000035a je           LBB0_57
	0x4d, 0x29, 0xfc, //0x00000360 subq         %r15, %r12
	0x49, 0x8d, 0x44, 0x24, 0xff, //0x00000363 leaq         $-1(%r12), %rax
	0x49, 0x39, 0xc6, //0x00000368 cmpq         %rax, %r14
	0x0f, 0x84, 0x3c, 0x00, 0x00, 0x00, //0x0000036b je           LBB0_48
	0x49, 0x39, 0xc2, //0x00000371 cmpq         %rax, %r10
	0x0f, 0x84, 0x33, 0x00, 0x00, 0x00, //0x00000374 je           LBB0_48
	0x49, 0x39, 0xc3, //0x0000037a cmpq         %rax, %r11
	0x0f, 0x84, 0x2a, 0x00, 0x00, 0x00, //0x0000037d je           LBB0_48
	0x4d, 0x85, 0xdb, //0x00000383 testq        %r11, %r11
	0x0f, 0x8e, 0x35, 0x00, 0x00, 0x00, //0x00000386 jle          LBB0_52
	0x49, 0x8d, 0x43, 0xff, //0x0000038c leaq         $-1(%r11), %rax
	0x49, 0x39, 0xc6, //0x00000390 cmpq         %rax, %r14
	0x0f, 0x84, 0x28, 0x00, 0x00, 0x00, //0x00000393 je           LBB0_52
	0x49, 0xf7, 0xd3, //0x00000399 notq         %r11
	0x4c, 0x89, 0xdf, //0x0000039c movq         %r11, %rdi
	0x48, 0x85, 0xff, //0x0000039f testq        %rdi, %rdi
	0x0f, 0x89, 0x98, 0x00, 0x00, 0x00, //0x000003a2 jns          LBB0_58
	0xe9, 0xb1, 0x00, 0x00, 0x00, //0x000003a8 jmp          LBB0_57
	//0x000003ad LBB0_48
	0x49, 0xf7, 0xdc, //0x000003ad negq         %r12
	0x4c, 0x89, 0xe7, //0x000003b0 movq         %r12, %rdi
	0x48, 0x85, 0xff, //0x000003b3 testq        %rdi, %rdi
	0x0f, 0x89, 0x84, 0x00, 0x00, 0x00, //0x000003b6 jns          LBB0_58
	0xe9, 0x9d, 0x00, 0x00, 0x00, //0x000003bc jmp          LBB0_57
	//0x000003c1 LBB0_52
	0x4c, 0x89, 0xd0, //0x000003c1 movq         %r10, %rax
	0x4c, 0x09, 0xf0, //0x000003c4 orq          %r14, %rax
	0x4d, 0x39, 0xf2, //0x000003c7 cmpq         %r14, %r10
	0x0f, 0x8c, 0x1d, 0x00, 0x00, 0x00, //0x000003ca jl           LBB0_55
	0x48, 0x85, 0xc0, //0x000003d0 testq        %rax, %rax
	0x0f, 0x88, 0x14, 0x00, 0x00, 0x00, //0x000003d3 js           LBB0_55
	0x49, 0xf7, 0xd2, //0x000003d9 notq         %r10
	0x4c, 0x89, 0xd7, //0x000003dc movq         %r10, %rdi
	0x48, 0x85, 0xff, //0x000003df testq        %rdi, %rdi
	0x0f, 0x89, 0x58, 0x00, 0x00, 0x00, //0x000003e2 jns          LBB0_58
	0xe9, 0x71, 0x00, 0x00, 0x00, //0x000003e8 jmp          LBB0_57
	//0x000003ed LBB0_55
	0x48, 0x85, 0xc0, //0x000003ed testq        %rax, %rax
	0x49, 0x8d, 0x46, 0xff, //0x000003f0 leaq         $-1(%r14), %rax
	0x49, 0xf7, 0xd6, //0x000003f4 notq         %r14
	0x4d, 0x0f, 0x48, 0xf4, //0x000003f7 cmovsq       %r12, %r14
	0x49, 0x39, 0xc2, //0x000003fb cmpq         %rax, %r10
	0x4d, 0x0f, 0x45, 0xf4, //0x000003fe cmovneq      %r12, %r14
	0x4c, 0x89, 0xf7, //0x00000402 movq         %r14, %rdi
	0x48, 0x85, 0xff, //0x00000405 testq        %rdi, %rdi
	0x0f, 0x89, 0x32, 0x00, 0x00, 0x00, //0x00000408 jns          LBB0_58
	0xe9, 0x4b, 0x00, 0x00, 0x00, //0x0000040e jmp          LBB0_57
	//0x00000413 LBB0_17
	0x4c, 0x29, 0xfb, //0x00000413 subq         %r15, %rbx
	0x0f, 0xbc, 0xfa, //0x00000416 bsfl         %edx, %edi
	0xe9, 0x07, 0x00, 0x00, 0x00, //0x00000419 jmp          LBB0_14
	//0x0000041e LBB0_13
	0x4c, 0x29, 0xfb, //0x0000041e subq         %r15, %rbx
	0x41, 0x0f, 0xbc, 0xf9, //0x00000421 bsfl         %r9d, %edi
	//0x00000425 LBB0_14
	0x48, 0x01, 0xdf, //0x00000425 addq         %rbx, %rdi
	//0x00000428 LBB0_15
	0x48, 0xf7, 0xd7, //0x00000428 notq         %rdi
	0x48, 0x8b, 0x75, 0xc0, //0x0000042b movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x5d, 0xc8, //0x0000042f movq         $-56(%rbp), %rbx
	0x48, 0x8b, 0x55, 0xd0, //0x00000433 movq         $-48(%rbp), %rdx
	0x48, 0x85, 0xff, //0x00000437 testq        %rdi, %rdi
	0x0f, 0x88, 0x1e, 0x00, 0x00, 0x00, //0x0000043a js           LBB0_57
	//0x00000440 LBB0_58
	0x49, 0x01, 0xff, //0x00000440 addq         %rdi, %r15
	0x48, 0x89, 0xd0, //0x00000443 movq         %rdx, %rax
	0xe9, 0x20, 0x00, 0x00, 0x00, //0x00000446 jmp          LBB0_59
	//0x0000044b LBB0_41
	0x49, 0x89, 0xc4, //0x0000044b movq         %rax, %r12
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x0000044e movq         $-1, %rdi
	0x4d, 0x85, 0xf6, //0x00000455 testq        %r14, %r14
	0x0f, 0x85, 0xec, 0xfe, 0xff, 0xff, //0x00000458 jne          LBB0_43
	//0x0000045e LBB0_57
	0x48, 0xf7, 0xd7, //0x0000045e notq         %rdi
	0x49, 0x01, 0xff, //0x00000461 addq         %rdi, %r15
	0x48, 0xc7, 0xc0, 0xfe, 0xff, 0xff, 0xff, //0x00000464 movq         $-2, %rax
	//0x0000046b LBB0_59
	0x49, 0x29, 0xdf, //0x0000046b subq         %rbx, %r15
	0x4c, 0x89, 0x3e, //0x0000046e movq         %r15, (%rsi)
	0x48, 0x83, 0xc4, 0x18, //0x00000471 addq         $24, %rsp
	0x5b, //0x00000475 popq         %rbx
	0x41, 0x5c, //0x00000476 popq         %r12
	0x41, 0x5d, //0x00000478 popq         %r13
	0x41, 0x5e, //0x0000047a popq         %r14
	0x41, 0x5f, //0x0000047c popq         %r15
	0x5d, //0x0000047e popq         %rbp
	0xc3, //0x0000047f retq         
	//0x00000480 LBB0_61
	0x48, 0xf7, 0xdf, //0x00000480 negq         %rdi
	0x48, 0x8b, 0x55, 0xd0, //0x00000483 movq         $-48(%rbp), %rdx
	0x48, 0x85, 0xff, //0x00000487 testq        %rdi, %rdi
	0x0f, 0x89, 0xb0, 0xff, 0xff, 0xff, //0x0000048a jns          LBB0_58
	0xe9, 0xc9, 0xff, 0xff, 0xff, //0x00000490 jmp          LBB0_57
	//0x00000495 LBB0_1
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000495 movq         $-1, %rax
	0xe9, 0xca, 0xff, 0xff, 0xff, //0x0000049c jmp          LBB0_59
	//0x000004a1 LBB0_8
	0x49, 0xc7, 0xc2, 0xff, 0xff, 0xff, 0xff, //0x000004a1 movq         $-1, %r10
	0x4d, 0x89, 0xfc, //0x000004a8 movq         %r15, %r12
	0x4d, 0x89, 0xe8, //0x000004ab movq         %r13, %r8
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x000004ae movq         $-1, %r14
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x000004b5 movq         $-1, %r11
	0xe9, 0xdc, 0xfd, 0xff, 0xff, //0x000004bc jmp          LBB0_31
	0x90, 0x90, 0x90, //0x000004c1 .p2align 2, 0x90
	// // .set L0_0_set_38, LBB0_38-LJTI0_0
	// // .set L0_0_set_42, LBB0_42-LJTI0_0
	// // .set L0_0_set_34, LBB0_34-LJTI0_0
	// // .set L0_0_set_40, LBB0_40-LJTI0_0
	// // .set L0_0_set_36, LBB0_36-LJTI0_0
	//0x000004c4 LJTI0_0
	0x17, 0xfe, 0xff, 0xff, //0x000004c4 .long L0_0_set_38
	0x76, 0xfe, 0xff, 0xff, //0x000004c8 .long L0_0_set_42
	0x17, 0xfe, 0xff, 0xff, //0x000004cc .long L0_0_set_38
	0x4d, 0xfe, 0xff, 0xff, //0x000004d0 .long L0_0_set_34
	0x76, 0xfe, 0xff, 0xff, //0x000004d4 .long L0_0_set_42
	0xec, 0xfd, 0xff, 0xff, //0x000004d8 .long L0_0_set_40
	0xec, 0xfd, 0xff, 0xff, //0x000004dc .long L0_0_set_40
	0xec, 0xfd, 0xff, 0xff, //0x000004e0 .long L0_0_set_40
	0xec, 0xfd, 0xff, 0xff, //0x000004e4 .long L0_0_set_40
	0xec, 0xfd, 0xff, 0xff, //0x000004e8 .long L0_0_set_40
	0xec, 0xfd, 0xff, 0xff, //0x000004ec .long L0_0_set_40
	0xec, 0xfd, 0xff, 0xff, //0x000004f0 .long L0_0_set_40
	0xec, 0xfd, 0xff, 0xff, //0x000004f4 .long L0_0_set_40
	0xec, 0xfd, 0xff, 0xff, //0x000004f8 .long L0_0_set_40
	0xec, 0xfd, 0xff, 0xff, //0x000004fc .long L0_0_set_40
	0x76, 0xfe, 0xff, 0xff, //0x00000500 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000504 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000508 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x0000050c .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000510 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000514 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000518 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x0000051c .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000520 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000524 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000528 .long L0_0_set_42
	0x32, 0xfe, 0xff, 0xff, //0x0000052c .long L0_0_set_36
	0x76, 0xfe, 0xff, 0xff, //0x00000530 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000534 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000538 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x0000053c .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000540 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000544 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000548 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x0000054c .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000550 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000554 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000558 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x0000055c .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000560 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000564 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000568 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x0000056c .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000570 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000574 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000578 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x0000057c .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000580 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000584 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000588 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x0000058c .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000590 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000594 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x00000598 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x0000059c .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x000005a0 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x000005a4 .long L0_0_set_42
	0x76, 0xfe, 0xff, 0xff, //0x000005a8 .long L0_0_set_42
	0x32, 0xfe, 0xff, 0xff, //0x000005ac .long L0_0_set_36
	//0x000005b0 .p2align 2, 0x00
	//0x000005b0 _MASK_USE_NUMBER
	0x02, 0x00, 0x00, 0x00, //0x000005b0 .long 2
}
 
