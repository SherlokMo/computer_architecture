MOV Ax, 5
MOV Bx, 1


L1:  
    CMP Ax, Bx
    JE EXIT
        INC Bx 
        LOOP L1
              
              
EXIT:
    HLT