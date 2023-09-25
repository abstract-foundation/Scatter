## Blocks on the Network
Blocks are how messages are distributed over the network. A node's job is 
to bundle messages into a block until it has reached the minimum block size and then
propagate the block throughout the network. A block has a minimum size of 1 megabyte.
A block has a maximum size of 10 megabytes. Meaning the network can scale for when 
demand is high. When a block is proposed each message must have the correct sequence for
the block to be valid. Once a block is validated on the network. The creator of the block is
then in charge of sending out the block to the computational nodes and storage nodes and creating
the appropiate block for the result of those actions.