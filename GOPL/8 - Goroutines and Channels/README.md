# **Chapter 8 - Goroutines and Channels**

## **8.1. Goroutines**

### **Overview**

<p style="text-align:justify">
&nbsp;&nbsp;&nbsp;&nbsp;Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought of as <u>light weight threads</u>. The cost of creating a Goroutine is tiny when compared to a thread. Hence its common for Go applications to have thousands of Goroutines running concurrently.</p>
<p style="text-align:justify">
&nbsp;&nbsp;&nbsp;&nbsp;Every program contains at least a single Goroutine and that Goroutine is known as the <b>main Goroutine</b>. All the Goroutines are working under the main Goroutines if the main Goroutine terminated, then all the goroutine present in the program also terminated. Goroutine always works in the background.
</p>

### **Goroutine vs Thread**
<table width=100%;>
    <tr>
        <th>Goroutine</th>
        <th>Thread</th>
    </tr>
    <tr>
        <td>Goroutines are managed by the go runtime.</td>
        <td>Operating system threads are managed by kernal.</td>
    </tr>
    <tr>
        <td>Goroutine are not hardware dependent.</td>
        <td>Threads are hardware dependent.</td>
    </tr>
    <tr>
        <td>Goroutines have easy communication medium known as channel.</td>
        <td>Thread doesnot have easy communication medium.</td>
    </tr>
    <tr>
        <td>Due to the presence of channel one goroutine can communicate with other goroutine with low latency.</td>
        <td>Due to lack of easy communication medium inter-threads communicate takes place with high latency.</td>
    </tr>
    <tr>
        <td>Goroutine doesnot have ID because go doesnot have Thread Local Storage.</td>
        <td>Threads have their own unique ID because they have Thread Local Storage.</td>
    </tr>
    <tr>
        <td>Goroutines are cheaper than threads.</td>
        <td>...</td>
    </tr>
    <tr>
        <td>Goroutines are cooperatively scheduled.</td>
        <td>Threads are preemptively scheduled.</td>
    </tr>
    <tr>
        <td>Goroutines have fasted startup time than threads.</td>
        <td>...</td>
    </tr>
    <tr>
        <td>Goroutine has growable segmented stacks.</td>
        <td>...</td>
    </tr>
</table>

### **Process vs Thread**
<table width=100%;>
    <tr>
        <th width=20%;>Comparison Basis</th>
        <th>Process</th>
        <th>Thread</th>
    </tr>
    <tr>
        <td>Definition</td>
        <td>A process is a program under executio</td>
        <td>A thread is a lightweight process that can be managed independently by a scheduler.</td>
    </tr>
    <tr>
        <td>Context switching time</td>
        <td>Processes require more time for context switching as they are more heavy.</td>
        <td>Threads require less time for context switching as they are lighter than processes.</td>
    </tr>
    <tr>
        <td>Memory Sharing</td>
        <td>Processes are totally independent and don’t share memory.</td>
        <td>A thread may share some memory with its peer threads.</td>
    </tr>
    <tr>
        <td>Communication</td>
        <td>Communication between processes requires more time than between threads.</td>
        <td>Communication between threads requires less time than between processes .</td>
    </tr>
    <tr>
        <td>Blocked</td>
        <td>If a process gets blocked, remaining processes can continue execution.</td>
        <td>If a user level thread gets blocked, all of its peer threads also get blocked.</td>
    </tr>
    <tr>
        <td>Resource Consumption</td>
        <td>Processes require more resources than threads.</td>
        <td>Threads generally need less resources than processes.</td>
    </tr>
    <tr>
        <td>Dependency</td>
        <td>Individual processes are independent of each other.</td>
        <td>Threads are parts of a process and so are dependent.</td>
    </tr>
    <tr>
        <td>Data and Code sharing</td>
        <td>Processes have independent data and code segments.</td>
        <td>A thread shares the data segment, code segment, files etc. with its peer threads.</td>
    </tr>
    <tr>
        <td>Treatment by OS</td>
        <td>All the different processes are treated separately by the operating system.</td>
        <td>All user level peer threads are treated as a single task by the operating system.</td>
    </tr>
    <tr>
        <td>Time for creation</td>
        <td>Processes require more time for creation.</td>
        <td>Threads require less time for creation.</td>
    </tr>
    <tr>
        <td>Time for termination</td>
        <td>Processes require more time for termination.</td>
        <td>Threads require less time for termination.</td>
    </tr>
</table>

### **Advantages of Goroutines over Thread**

- <p style="text-align:justify">Goroutines are cheaper than threads.</p>
- <p style="text-align:justify">Goroutine are stored in the stack and the size of the stack can grow and shrink according to the requirement of the program. But in threads, the size of the stack is fixed.</p>
- <p style="text-align:justify">Goroutines can communicate using <b>the channel</b> and these channels are specially designed to prevent <b>race conditions</b> when accessing shared memory using Goroutines. Channels can be thought of as <b>a pipe</b> using which Goroutines communicate.</p>
- <p style="text-align:justify">Suppose a program has one thread, and that thread has many Goroutines associated with it. If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread. All these details are hidden from the programmers.</p>

### **Channels**
<p style="text-align:justify;"></p>

### **Deadlock**

### **Cancellation**

### **What is Race Condition?**

### **Mutex**

### **Concurrency Patterns**

### **fan-in & fan-out**
<p style="text-align:justify;">fan-in is a multiplexing strategy where the inputs of several channels are combined to produce an output channel. fan-out is demultiplexing strategy where a single channel is split into multiple channels.</p>