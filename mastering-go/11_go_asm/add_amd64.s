// see https://davidwong.fr/goasm/add
// see https://habr.com/ru/post/349384/
TEXT ·add(SB),$0-24
    MOVQ x+0(FP), BX   
    MOVQ y+8(FP), BP
    ADDQ BP, BX
    MOVQ BX, ret+16(FP)
    RET
