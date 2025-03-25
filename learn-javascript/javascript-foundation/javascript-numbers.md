# JavaScript Numbers

JavaScript has only one type of number. Numbers can be written with or without decimals.

```javascript
let x = 3.14; // A number with decimals
let y = 3; // A number without decimals
```

Extra large or extra small numbers can be written with scientific (exponent) notation:

```javascript
let x = 123e5; // 12300000
let y = 123e-5; // 0.00123
```

## JavaScript Numbers are Always 64-bit Floating Point

Unlike many other programming languages, JavaScript does not define different types of numbers, like integers, short, long, floating-point etc.

JavaScript numbers are always stored as double precision floating point numbers, following the international IEEE 754 standard.

This format stores numbers in 64 bits, where the number (the fraction) is stored in bits 0 to 51, the exponent in bits 52 to 62, and the sign in bit 63:

| Value            | Exponent          | Sign       |
| ---------------- | ----------------- | ---------- |
| 52 bits (0 - 51) | 11 bits (52 - 62) | 1 bit (63) |
