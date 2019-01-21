# CHALLANGE SHELL


GOLANG,  Go Routines, Channels

The assignment is a Clock.

TICK01 - "tick" every second
TICK02 - "tock" every minute
TICK03 - "bong" every hour
BOOM - "Kaboooom" at the end of the Clock

The Tick Messages can be changed at random times.

Almost all of the files are using a SHELL.json file for
configuration. 
The operator can edit and save the SHELL.json file to change the messages. The Messges will change after the 
TICK03 event.

I approach the system in several strategies.

NOTE:
1. I should change the CONFIG Event to its own ticker.
2. I did not add _test suite to this project as I do not know how to do _test on something that fires go routines.



# V01
This is a minimalist app.  
However, does not respect the Single Message per second rule.  I can fix this but I did not notice the rule until 
just now.



# V02
I tried to handle the Single Tick per second.
But I am really tired.  After the wedding weekend, I over slept. Not sure this is working.


# V03
I am still working on this..
but it will...
1. Act as an ON-PREM IoT device in a minimalist Scratch Docker container.
2. TICK events messages are Sent to Cloud REpositiory ( Firebase )
3. Cloud relays to iPHONE and ANDROID devices.
4. Mobile devices can change MESSAGE on TICK Event remotely.


# V04
< stay tune >
Add a vue.js portal that will update on TICK events relayed from Firebase.

# V10
< stay tune >
totally different solution
1. Using GoLang's Microservice framework.
2. Attempt to do redundant servers with one master and failovers.
3. Messages would be transmitted over ETCD
4. Clients would receive TICK messages over ETCD
5. Client would modify Tick Message over ETCD
I think this is a really interesting strategy for a controller on a remote/ or low connectivity platform where battery size and life are high priorities.




