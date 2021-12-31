MOV Ax, 5h
MOV Bx, 6h

CMP Ax, Bx

JA L1
    DEC Ax
    JMP EXIt
L1:
    INC Bx
EXIT:
    HLT