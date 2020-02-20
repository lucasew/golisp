package lex

type LexByte byte

func (b LexByte) IsByte(ob byte) bool {
    return b == LexByte(ob)
}

func (b LexByte) IsByteNumber() bool {
    return b >= '0' && b <= '9'
}

func (b LexByte) IsByteBigLetter() bool {
    return b >= 'A' && b <= 'Z'
}

func (b LexByte) IsByteSmallLetter() bool {
    return b >= 'a' && b <= 'z'
}

func (b LexByte) IsByteSpecialSymbol() bool {
    return b.IsByte('+') || b.IsByte('-') || b.IsByte('/') || b.IsByte('*') || b.IsByte('&') || b.IsByte('|') || b.IsByte('%') || b.IsByte('.') // TODO: Add the rest of
}

func (b LexByte) IsByteUnderline() bool {
    return b.IsByte('_')
}

func (b LexByte) IsByteLetter() bool {
    return b.IsByteSmallLetter() || b.IsByteBigLetter()
}

func (b LexByte) IsByteE() bool {
    return b.IsByte('e') || b.IsByte('E')
}

func (b LexByte) IsByteColon() bool {
    return b.IsByte(':')
}

func (b LexByte) IsOpenPar() bool {
    return b.IsByte('(')
}

func (b LexByte) IsClosePar() bool {
    return b.IsByte(')')
}

func (b LexByte) IsBlank() bool {
    return b.IsByte(' ') || b.IsByte('\n') || b.IsByte('\r')
}

func (b LexByte) IsStringMark() bool {
    return b.IsByte('"')
}

func (b LexByte) IsHash() bool {
    return b.IsByte('#')
}

func (b LexByte) IsDot() bool {
    return b.IsByte('.')
}

func (b LexByte) IsSlash() bool {
    return b.IsByte('/')
}

func (b LexByte) IsHexadecimal() bool {
    if (b >= '0' && b <= '9') {
        return true
    }
    if (b >= 'a' && b <= 'f') {
        return true
    }
    if (b >= 'A' && b <= 'F') {
        return true
    }
    return false
}
