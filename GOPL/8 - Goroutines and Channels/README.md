# **Chapter 8 - Goroutines and Channels**

## **8.1. Goroutines**

### **Overview**

<p style="text-align:justify">
&nbsp;&nbsp;&nbsp;&nbsp;Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought of as <u>light weight threads</u>. The cost of creating a Goroutine is tiny when compared to a thread. Hence its common for Go applications to have thousands of Goroutines running concurrently.</p>
<p style="text-align:justify">
&nbsp;&nbsp;&nbsp;&nbsp;Every program contains at least a single Goroutine and that Goroutine is known as the <b>main Goroutine</b>. All the Goroutines are working under the main Goroutines if the main Goroutine terminated, then all the goroutine present in the program also terminated. Goroutine always works in the background.
</p>

### **Goroutine vs Thread**
<table>
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

### **Advantages of Goroutines over Thread**

- <p style="text-align:justify">Goroutines are cheaper than threads.</p>
- <p style="text-align:justify">Goroutine are stored in the stack and the size of the stack can grow and shrink according to the requirement of the program. But in threads, the size of the stack is fixed.</p>
- <p style="text-align:justify">Goroutines can communicate using <b>the channel</b> and these channels are specially designed to prevent <b>race conditions</b> when accessing shared memory using Goroutines. Channels can be thought of as <b>a pipe</b> using which Goroutines communicate.</p>
- <p style="text-align:justify">Suppose a program has one thread, and that thread has many Goroutines associated with it. If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread. All these details are hidden from the programmers.</p>

### **Channels**

### **Cancellation**

### **What is Race Condition?**

### **Mutex**

