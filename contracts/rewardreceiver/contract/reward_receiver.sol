pragma solidity ^0.5.10;

contract RewardReceiver {
    function() payable external{}
    function sendTo(address payable receiver,uint256 amount) public {
        receiver.transfer(amount);
    }
}
