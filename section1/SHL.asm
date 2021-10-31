; SHL is logical shift, which fills newely created bit \n
; position with zero 

; Shifting left n bits multiplies the operand with 2^n

MOV Al, 5 ; 0000 0101 = 5
SHL Al, 1 ; 0000 1010 = 10