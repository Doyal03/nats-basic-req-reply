# nats-basic-req-reply
The ReplyServer is the listening server, and the RequestServer is the one making the requests.

The RequestServer starts sending the message one by one; if the ReplyServer is not up, the messages will be lost.

To test the working :-
Clone the two folders and launch two different cmd prompts and respectively run "go run ." on both
![image](https://user-images.githubusercontent.com/91259249/191293337-9ff2dc93-d140-4397-9f10-a22c268d9b77.png)

IMP:- launch ReplyServer first to avoid missing messages!

![image](https://user-images.githubusercontent.com/91259249/191293842-2c50b3a9-0176-418f-bdb3-808fcafd25b7.png)

Once all messages are sent from RequestServer, it will display "Message delivered". Its not checking in any way whether the messages were actually consumed by the ReplyServer or not.
