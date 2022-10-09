def caesarCipher(s, k):
    k %= 26
    lower = "abcdefghijklmnopqrstuvwxyz"
    upper = lower.upper()
    key = dict(zip(lower, lower[k:] + lower[:k])) | dict(zip(upper, upper[k:] + upper[:k]))
    return s.translate(s.maketrans(key))