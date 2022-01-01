MOV Al, 05h

MOV Bl, Al

L1: 
    DEC Al
    CMP Al, 00h
    JE EXIT    
        MUL Bl, Al 
        LOOP L1
EXIT:  
    HLT