MOV [500], 5
MOV [501], 40
MOV [502], 30
MOV [503], 50
MOV [504], 20
MOV [505], 10

MOV SI, 500

MOV Cl, [SI]

INC SI
MOV Al, [SI]  

L1:
    DEC Cl
    CMP Cl, 00h
    JE EXIT
        INC SI
        CMP Al, [SI]
        JL L2
        L2:
            MOV Al, [SI] 
        LOOP L1
EXIT:
    MOV [506], Al
    HLT