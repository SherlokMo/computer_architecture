MOV Ax, 6
MOV Bx, -6

CMP Ax, Bx

JG L1
    DEC Ax
    JMP EXIT

L1:
    INC Bx          
 
EXIT:
    HLT

