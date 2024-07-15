# grpc-ticket-service

#datastore, package is to manage all the data related to 

  -> users
  -> trains
  -> tickets
  -> allocations

  test cases are added in the go file, datastore_test for these

#ticketservice package the grpc related code and the server struct

quoting the details from task here for easy reference. refer server.go for the interface implemenations for each of the grpc APIs.
implementation of client for testing is added in testClients/client.go and allocation.go

#1---> Create API where you can submit a purchase for a ticket.  Details included in the receipt are: 
		a. From, To, User , price paid.
	       i. User should include first and last name, email address
		The user is allocated a seat in the train.  Assume the train has only 2 sections, section A and section B.
   API 1: TicketPurchaseService

  
#2 ---> An API that shows the details of the receipt for the user

   API2 : GetReceipt


#3 ---> An API that lets you view the users and seat they are allocated by the requested section
   API3 : AllocationStatus

#4 --> An API to remove a user from the train
    CancelTicket
#5. An API to modify a user's seat

    SeatReallocate
#6. GetAllTickets(DummyMessage) returns (TicketsMinListRes) ; is added for testing purpose


Test Coverage for datastore alone is added. Remain has to be explored how to do it.

vijayakanth@msf-HP-348-G7:~/grpc/ticket_service$ go test ./datastore/ -cover
ok      github.com/vijayakanth89/grpc-ticket-service/datastore  (cached)        coverage: 88.2% of statements
