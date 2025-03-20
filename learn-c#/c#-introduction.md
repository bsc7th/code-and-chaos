# Learn C#

# Common Principles

1. **KISS** - Keep It Simple Stupid
2. **DAMP** - Descriptive And Meaningfup Phrases
3. **WET** - Write Everything Twice
4. **DRY** - Do Not Repeat Yourself

# Datatype Categories

1. Value Type

- Definition: A value type directly holds data. When a value type is assigned to a new variable, a copy of the value is created.
- Examples: All primitive data types like `int`, `char`, `bool`, `float`, `double`, `byte`, `decimal`, `float`, `long`, `sbyte`, `short`, `uint`, `ulong`, `ushort` as well as structs (struct).

- `int` - represents 32-bit signed integer type
- `char` - represents 16-bit Unicode Character - range from `U +0000` to `U +ffff`. The Default Value is `\0'`.
- `bool` - represents boolean value - `true` or `false`. The Default Value is `false`.

- Storage: The data is stored on the stack (which is a region in memory that is fast and automatically managed).

**Key Characteristics:**

- When you assign one value type to another, the data is copied.
- Changes to one variable do not affect others, because they each hold their own copy of the data.
- Value types have no reference to any other memory location.

Example:

```csharp
int x = 10;
int y = x;  // y is a copy of x
y = 20;     // Changing y doesn't affect x
```

2. Reference Type

- Definition: A reference type stores a reference (or address) to the actual data rather than the data itself. When a reference type is assigned to a new variable, both variables refer to the same object.
- Examples: Objects, arrays, strings, and classes.
- Storage: The reference itself is stored on the stack, but the actual data (object) is stored in the heap (a memory area managed by the runtime for dynamic memory allocation).

**Key Characteristics:**

- When you assign one reference type to another, both variables point to the same data in memory.
- Changes made through one reference will affect the other, because they both refer to the same object.
- Reference types involve garbage collection to automatically manage memory allocation and deallocation.

Example:

```csharp
class Person
{
    public string Name;
}

Person p1 = new Person();
p1.Name = "Alice";

Person p2 = p1;  // p2 refers to the same object as p1
p2.Name = "Bob";  // Changing the Name property of p2 also affects p1

Console.WriteLine(p1.Name);  // Output: Bob
```

3. Pointer Type (Note: Not used in C# typically)

- Definition: A pointer type stores the memory address of another variable. Pointers allow direct memory access and manipulation, which is a feature of languages like C and C++.
- C# and Pointers: C# does not use pointers in everyday programming, and you generally don’t need them. They can only be used in unsafe code, a special feature for advanced scenarios where direct memory manipulation is necessary.
- In unsafe code (which needs explicit permission and the unsafe keyword), you can use pointers in C#, similar to C++.
- However, this is considered unsafe and should only be used when absolutely necessary, as it can lead to memory management issues like leaks or corruption.

Example (Unsafe Code):

```csharp
unsafe
{
    int num = 10;
    int* p = &num;  // p is a pointer to num
    Console.WriteLine(*p);  // Dereferencing p to get the value of num
}
```

**Note:** You must enable unsafe code in your project settings to run this.

Summary:

- Value Type: Stores actual data; copies the value when assigned to another variable.
- Reference Type: Stores a reference (address) to the data; both variables refer to the same object in memory.
- Pointer Type: Directly stores a memory address; mostly unused in C#, as it requires "unsafe" code.
