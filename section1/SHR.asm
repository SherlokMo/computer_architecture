; SHR is logical shift, which fills newely created bit \n
; position with zero 

; Shifting right n bits divides the operand with 2^n

MOV Al, 80 ; 1000 0000 = 80
SHL Al, 1  ; 0100 0000 = 40
SHL Al, 2  ; 0001 0000 = 10
