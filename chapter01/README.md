# **Table of Content**
- [**Table of Content**](#table-of-content)
- [**Chapter 1: Data Structures and Algorithms**](#chapter-1-data-structures-and-algorithms)
  - [**Classification of data structures**](#classification-of-data-structures)
    - [**Lists**](#lists)
    - [**Tuples**](#tuples)
    - [**Heaps**](#heaps)
  - [**Structural design patterns**](#structural-design-patterns)
    - [**Adapter**](#adapter)
      - [**Example I**](#example-i)
      - [**Example II**](#example-ii)
    - [Bridge](#bridge)
      - [**Example I**](#example-i-1)
      - [**Example II**](#example-ii-1)
# **Chapter 1: Data Structures and Algorithms**

A data structure is the organization of data to reduce the storage space used and to reduce the difficulty
while performing different tasks. Data structures are used to handle and work with large amounts of data in
various fields, such as database management and internet indexing services.

In this chapter, we will focus on the definition of abstract datatypes, classifying data structures into linear, nonlinear, homogeneous, heterogeneous, and dynamic types.Abstract datatypes, such as Container, List, Set, Map,
Graph, Stack, and Queue, are presented in this chapter. We will also cover the performance analysis of data
structures, choosing the right data structures, and structural design patterns.

The reader can start writing basic algorithms using the right data structures in Go. Given a problem, choosing
the data structure and different algorithms will be the first step. After this, doing performance analysis is
the next step. Time and space analysis for different algorithms helps compare them and helps you choose the
optimal one. It is essential to have basic knowledge of Go to get started.

> In this chapter, we will cover the following topics:
>
> - Classification of data structures and structural design patterns
> - Representation of algorithms
> - Complexity and performance analysis
> - Brute force algorithms
> - Divide and conquer algorithms
> - Backtracking algorithms
>

## **Classification of data structures**

The term **data structure** refers to the organization of data in a computer's memory, in order
to retrieve it quickly for processing. It is a scheme for data organization to decouple the
functional definition of a data structure from its implementation. A data structure is chosen
based on the problem type and the operations performed on the data.  

| **Linear** | **Non-Linear** | **Homogeneous**  | **Hetrogeneous**  | **Dynamic**    |
|------------|----------------|------------------|-------------------|----------------|
| - Lists    | - tress        | - 2D-Arrays      | - Linked Lists    | - Dictionaries |
| - Sets     | - Tables       | - Multi-D-Arrays | - Ordered Lists   | - Tree Sets    |
| - Tuples   | - Containers   |                  | - Unordered Lists | - Sequences    |
| - Queues   |                |                  |                   |                |
| - Stacks   |                |                  |                   |                |
| - Heaps    |                |                  |                   |                |

### **Lists**

A list is a sequence of elements. Each element can be connected to another with a link in a
forward or backward direction. The element can have other payload properties. This data
structure is a basic type of container. Lists have a variable length and developer can remove
or add elements more easily than an array. Data items within a list need not be contiguous
in memory or on disk. Linked lists were proposed by Allen Newell, Cliff Shaw, and Herbert 
A. Simon at RAND Corporation.

Run the code 

```bash
cd chapter01 \
go run . list

```

### **Tuples**

A tuple is a finite sorted list of elements. It is a data structure that groups data. Tuples are
typically immutable sequential collections. The element has related fields of different
datatypes. The only way to modify a tuple is to change the fields. Operators such as + and *
can be applied to tuples. A database record is referred to as a tuple. In [this example](tuple.go), power series of integers are calculated and the square and cube of the integer is returned as a tuple. Run the code

```bash
cd chapter01 \
go run . tuple

```
### **Heaps**

A heap is a data structure that is based on the heap property. The heap data structure is
used in selection, graph, and k-way merge algorithms. Operations such as finding,
merging, insertion, key changes, and deleting are performed on heaps. Heaps are part of
the container/heap package in Go. According to the heap order (maximum heap)
property, the value stored at each node is greater than or equal to its children.

If the order is descending, it is referred to as a maximum heap; otherwise, it's a minimum
heap. The heap data structure was proposed by J.W.J. Williams in 1964 for a heap sorting
algorithm. It is not a sorted data structure, but partially ordered. [This example](heap.go)
shows how to use the container/heap package to create a heap data structure. Run the code

```bash
cd chapter01 \
go run . heap

```

## **Structural design patterns**

Structural design patterns describe the relationships between the entities. They are used to
form large structures using classes and objects. These patterns are used to create a system
with different system blocks in a flexible manner. Adapter, bridge, composite, decorator,
facade, flyweight, private class data, and proxy are the Gang of Four (GoF) structural
design patterns. The private class data design pattern is the other design pattern covered in
this section.
We will take a look at adapter and bridge design patterns in the next sections.

### **Adapter**

The adapter pattern provides a wrapper with an interface required by the API client to link
incompatible types and act as a translator between the two types. The adapter uses the
interface of a class to be a class with another compatible interface. When requirements
change, there are scenarios where class functionality needs to be changed because of
incompatible interfaces.

The **dependency inversion principle** can be adhered to by using the adapter pattern, when a
class defines its own interface to the next level module interface implemented by an
*adapter* class. **Delegation** is the other principle used by the adapter pattern. Multiple
formats handling source-to-destination transformations are the scenarios where the adapter
pattern is applied.

The adapter pattern comprises the *target*, *adaptee*, *adapter*, and *client*:

- Target is the interface that the client calls and invokes methods on the adapter
and adaptee.
- The client wants the incompatible interface implemented by the adapter.
- The adapter translates the incompatible interface of the adaptee into an interface
that the client wants.
#### **Example I**

Let's say you have an **IProcessor** interface with a *process* method, the **Adapter**
class implements the *process* method and has an **Adaptee** instance as an attribute.
The **Adaptee** class has a *convert* method and an *adapterType* instance variable. The
developer while using the API client calls the *process* interface method to invoke *convert*
on **Adaptee**. You can see the code [here](adapter.go). Run the code

```bash
cd chapter01 \
go run . adapter1

```
#### **Example II** 

See [reference](https://refactoring.guru/design-patterns/adapter/go/example#)

Assume we have a client code that expects some features of an object (Lightning port), but we have another object called adaptee (Windows laptop) which offers the same functionality but through a different interface (USB port). 

This is where the Adapter pattern comes into the picture. We create a struct type known as adapter that will:

- Adhere to the same interface which the client expects (Lightning port).

- Translate the request from the client to the adaptee in the form that the adaptee expects. The adapter accepts a Lightning connector and then translates its signals into a USB format
and passes them to the USB port in windows laptop.

See the codes [here](adapter). Run the code

```bash
cd chapter01 \
go run . adapter2

```

### Bridge

Bridge decouples the implementation from the abstraction. The abstract base class can be
subclassed to provide different implementations and allow implementation details to be
modified easily. The interface, which is a bridge, helps in making the functionality of
concrete classes independent from the interface implementer classes. The bridge patterns
allow the implementation details to change at runtime.

The bridge pattern demonstrates the principle, preferring composition over inheritance.
It helps in situations where one should subclass multiple times orthogonal to each other.
Runtime binding of the application, mapping of orthogonal class hierarchies, and platform
independence implementation are the scenarios where the bridge pattern can be applied.

The bridge pattern components are *abstraction*, *refined abstraction*, *implementer*, and
*concrete implementer*. **Abstraction** is the interface implemented as an abstract class
that clients invoke with the method on the **concrete implementer**. Abstraction maintains
a has-a relationship with the implementation, instead of an is-a relationship. The has-a 
relationship is maintained by composition. Abstraction has a reference of the implementation.
**Refined abstraction** provides more variations than abstraction.


Bridge is a structural design pattern that divides business logic or huge class into separate
class hierarchies that can be developed independently.

One of these hierarchies (often called the Abstraction) will get a reference to an object of the
second hierarchy (Implementation). The abstraction will be able to delegate some (sometimes, most)
of its calls to the implementations object. Since all implementations will have a common interface,
they’d be interchangeable inside the abstraction. Let's try it.

#### **Example I**

Let's say ***IDrawShape*** is an interface with the *drawShape* method. **DrawShape** implements
the ***IDrawShape*** interface. We create an ***IContour*** bridge interface with the *drawContour*
method. The **contour** class implements the ***IContour*** interface. The **ellipse** class will
have a, b , r properties and **drawShape** (an instance of **DrawShape**). The **ellipse** class 
implements the **contour** bridge interface to implement the *drawContour* method. The *drawContour*
method calls the *drawShape* method on the **drawShape** instance. Run the [code](bridge.go) 

```bash
cd chapter01 \
go run . bridge1

```
#### **Example II**

Say, you have two types of computers: Mac and Windows. Also, two types of printers: Epson and HP. Both computers and printers need to work with each other in any combination. The client doesn’t want to worry
about the details of connecting printers to computers.

If we introduce new printers, we don’t want our code to grow exponentially. Instead of creating four structs for the 2*2 combination, we create two hierarchies:

- Abstraction hierarchy: this will be our computers
- Implementation hierarchy: this will be our printers

These two hierarchies communicate with each other via a Bridge, where the Abstraction (computer) contains a reference to the Implementation (printer). Both the abstraction and implementation can be developed independently without affecting each other. You can run the [code](bridge) by the following command:

```bash
cd chapter01 \
go run . bridge2

```
