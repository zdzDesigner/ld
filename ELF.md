## ELF

ELF（Executable and Linkable Format）是一种可执行文件和可链接文件的标准文件格式。
[wiki](https://en.wikipedia.org/wiki/Executable_and_Linkable_Format)

### ELF Header

- ident

```flow
# 识别和描述ELF文件的标识符  

EI_MAG0         :魔数字节0
EI_MAG1         :魔数字节1
EI_MAG2         :魔数字节2
EI_MAG3         :魔数字节3
EI_CLASS        :文件类别,指示文件是01:32位还是02:64位
EI_DATA         :数据编码方式01:小端
EI_VERSION      :文件版本
EI_OSABI        :操作系统/ABI标识符
EI_ABIVERSION   :ABI版本
EI_PAD          :填充字节开始位置
EI_NIDENT       :文件标识符长度
```

- type

```flow
# 文件类型，指示该文件是可执行文件、共享库还是目标文件等。

ET_NONE (0): 未知类型的文件。
ET_REL (1): 可重定位文件，包含可重定位的目标代码和数据。
ET_EXEC (2): 可执行文件，包含可执行的目标代码。
ET_DYN (3): 共享对象文件，用于动态链接的共享库。
ET_CORE (4): 核心转储文件，记录了程序在崩溃时的状态。
ET_LOOS (0xfe00): 保留给操作系统特定语义的起始值。
ET_HIOS (0xfeff): 保留给操作系统特定语义的结束值。
ET_LOPROC (0xff00): 保留给处理器特定语义的起始值。
ET_HIPROC (0xffff): 保留给处理器特定语义的结束值。

```

- machine

```flow
# 目标体系结构，指示文件的目标硬件体系结构。

EM_NONE (0): 未指定体系结构。
EM_M32 (1): AT&T WE 32100。
EM_SPARC (2): SPARC。
EM_386 (3): Intel 80386。
EM_68K (4): Motorola 68000。
EM_88K (5): Motorola 88000。
EM_860 (7): Intel 80860。
EM_MIPS (8): MIPS R3000。
EM_PARISC (15): HP PA-RISC。
EM_SPARC32PLUS (18): SPARC v8+。
EM_PPC (20): PowerPC。
EM_ARM (40): ARM。
EM_X86_64 (62): AMD x86-64。
EM_AARCH64 (183): ARM 64-bit
```

- 目标体系结构
- 入口点地址等

_读取前 4 个字节 7f 45 4c 46_

```sh
➜ hexdump -C -n 8 out/init/out.o
00000000  7f 45 4c 46 02 01 01 00                           |.ELF....|
00000008
```

00000000000000000000000000000000
