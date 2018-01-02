![alt text][gooopher]

[gooopher]:https://raw.githubusercontent.com/ismayilmalik/golang-design-patterns/master/gopher.jpg "Gooopher.."

# Design patterns in golang
>A beginner guide... happy coding!

## Table of Contents

- [Creational patterns](#creational-patterns)
    - [Singleton](#singleton)
    - [Builder](#builder)
    - [Factory](#factory)
    - [Abstract Factory](#abstract-factory)
- [Structural patterns](#structural-patterns)
    - [Composition](#composition)
- [Behavioral patterns](#behavioral-patterns)
    - [Chain of responsibility](#chain-of-responsibility)
    - [Strategy](#strategy)
- [Concurrency patterns](#concurrency-patterns)

---

## Creational patterns
Creational design patterns abstract the instantiation process. They help make a system independent of how its objects are created, composed, and represented.

### Singleton
Ensure a class only has one instance, and provide a global point of access to it.

### Builder
Separate the construction of a complex object from its representation so that the
same construction process can create different representations.

### Factory
Define an interface for creating an object, but let subclasses decide which class to
instantiate. Factory Method lets a class defer instantiation to subclasses.

### Abstract Factory
Provide an interface for creating famili es of related or dependent objects without
specifying their concrete classes.

## Structural patterns
Structural patterns are concerned with how classes and objects are composed to form
larger structures. Structural class patterns use inheritance to compose interfaces or implementations.
As a simple example, consider how multiple inheritance mixes two or
more classes into one.

### Composition
Compose objects into tree structures to represent part-whole hierarchies. Composite
lets clients treat individual objects and compositions of objects uniformly.

## Behavioral patterns
Behavioral patterns are concerned with algorithms and the assignment of responsibilities
between objects. Behavio ral patterns describe not just patterns of objects or classes
but also the patterns of communication between them. These patterns characterize
complex control flow that's difficult to follow at run-time. They shift your focus away
from flow of control to let you concent ratejust on the way objects are interconnected.

### Chain of responsibility
Avoid coupling the sender of a request to its receiver by giving more than one
object a chance to handle the request. Chain the receiving objects and pass the
request along the chain until an object handles it.

### Strategy
Define a family oafl gorit h ms, encapsulate each one, and make them interchangeable.
Strategy lets the algorithm vary independently from clients that use it.

## Concurrency patterns
Pattenrs for concurrent work and parallel execution in Go.
