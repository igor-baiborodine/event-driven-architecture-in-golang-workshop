# Event-Driven Architecture in Golang

This workshop covers material from
the ["Event-Driven Architecture in Golang: Building complex systems with asynchronicity and eventual consistency"](https://www.amazon.ca/dp/1803238011/)
book. It contains the same source code as in the original [repository](https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang)
but is updated and executed against the latest versions of dependent Go packages.

**Table of Contents**

- [X] 1.Introduction to Event-driven Architectures 
- [X] 2.[Supporting Patterns in Brief](/Chapter02)
- [X] 3.[Design and Planning](/Chapter03) ([`65ef92b`](https://github.com/igor-baiborodine/event-driven-architecture-in-golang-workshop/commit/65ef92b7b499bea5b42768376b1b9dbddb890593))
- [X] 4.[Event Foundations](/Chapter04) ([`6a2d842`](https://github.com/igor-baiborodine/event-driven-architecture-in-golang-workshop/commit/6a2d84294cdb43c14252c4bef2035c5f85979196))
- [X] 5.[Tracking Changes with Event Sourcing](/Chapter05) ([`020069a`](https://github.com/igor-baiborodine/event-driven-architecture-in-golang-workshop/commit/020069ae1ab0472dd08d59747157ba8be935934a))
- [X] 6.[Asynchronous Connections](/Chapter06) ([`31cc914`](https://github.com/igor-baiborodine/event-driven-architecture-in-golang-workshop/commit/31cc91402392c3f193d980474bbd0bd4bbe1032d))
- [X] 7.[Event-Carried State Transfer](/Chapter07) ([`f4c47c1`](https://github.com/igor-baiborodine/event-driven-architecture-in-golang-workshop/commit/f4c47c10127080980c8de2839841ed149a263eb1))
- [X] 8.[Message Workflows](/Chapter08) ([`8edfc5b`](https://github.com/igor-baiborodine/event-driven-architecture-in-golang-workshop/commit/8edfc5bce342e4b854a171b520c34600a4749b3d))
- [X] 9.[Transactional Messaging](/Chapter09) ([`40e5d9f`](https://github.com/igor-baiborodine/event-driven-architecture-in-golang-workshop/commit/40e5d9ff754b486b642c2706d12db92c7e6ac261))
- [ ] 10.[Testing](/Chapter10)
- [ ] 11.Deploying Applications to the Cloud
- [ ] 12.Monitoring and Observability

<a href="https://www.packtpub.com/product/event-driven-architecture-in-golang/9781803238012?utm_source=github&utm_medium=repository&utm_campaign=9781803238012"><img src="https://static.packt-cdn.com/products/9781803238012/cover/smaller" alt="" height="256px" align="right"></a>

This is the code repository for [Event-Driven Architecture in Golang](https://www.packtpub.com/product/event-driven-architecture-in-golang/9781803238012?utm_source=github&utm_medium=repository&utm_campaign=9781803238012), published by Packt.

**Building complex systems with asynchronicity and eventual consistency**

## What is this book about?
Event-Driven Architecture in Golang is an approach used to develop applications that shares state changes asynchronously, internally, and externally using messages. EDA applications are better suited at handling situations that need to scale up quickly and the chances of individual component failures are less likely to bring your system crashing down. 

This book covers the following exciting features:
* Understand different event-driven patterns and best practices
* Plan and design your software architecture with ease
* Track changes and updates effectively using event sourcing
* Test and deploy your sample software application with ease
* Monitor and improve the performance of your software architecture

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1803238011) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png" 
alt="https://www.packtpub.com/" border="5" /></a>

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter02.

The code will look like the following:
```
BEGIN;
-- execute queries, updates, inserts, deletes ...
PREPARE TRANSACTION 'bfa1c57a-d99d-4d74-87a9-3aaabcc754ee';
```

**Following is what you need for this book:**
This hands-on book is for intermediate-level software architects, or senior software engineers working with Golang and interested in building asynchronous microservices using event sourcing, CQRS, and DDD. Intermediate-level knowledge of the Go syntax and concurrency features is necessary.

With the following software and hardware list you can run all code files present in the book (Chapter 1-12).
### Software and Hardware List
| Chapter | Software required | OS required |
| -------- | ------------------------------------ | ----------------------------------- |
| 1-12 | Go 1.18+ | Windows, Mac OS X, and Linux (Any) |
| 1-12 | Docker 20.10.x | Windows, Mac OS X, and Linux (Any) |
| 1-12 | NATS 2.4 | Windows, Mac OS X, and Linux (Any) |


We also provide a PDF file that has color images of the screenshots/diagrams used in this book. [Click here to download it](https://packt.link/qgf1O).

### Related products
* Go for DevOps 
[[Packt]](https://www.packtpub.com/product/go-for-devops/9781801818896?utm_source=github&utm_medium=repository&utm_campaign=9781801818896) [[Amazon]](https://www.amazon.com/dp/1801818894)

* Mastering Go - Third Edition 
[[Packt]](https://www.packtpub.com/product/mastering-go-third-edition/9781801079310?utm_source=github&utm_medium=repository&utm_campaign=9781801079310) [[Amazon]](https://www.amazon.com/dp/1801079315)


## Get to Know the Author
**Michael Stack**
is a solutions architect with more than 20 years experience developing software. He has been working with Golang for over seven years. During his career he has developed numerous N-tier applications, and is focused on development involving microservices and other distributed applications running in the cloud.

### Download a free PDF

 <i>If you have already purchased a print or Kindle version of this book, you can get a DRM-free PDF version at no cost.<br>Simply click on the link to claim your free PDF.</i>
<p align="center"> <a href="https://packt.link/free-ebook/9781803238012">https://packt.link/free-ebook/9781803238012 </a> </p>