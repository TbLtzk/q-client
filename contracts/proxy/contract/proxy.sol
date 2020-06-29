pragma solidity ^0.5.10;

interface Validators {
    function addValidator(address) external;

    function getValidatorsList() external view returns (address[] memory);
}

contract ValidatorsProxy {
    address ownerAddress;
    Validators delegateContract;


    function _isContract(address addr) internal view returns (bool) {
        uint size;
        assembly {size := extcodesize(addr)}
        return size > 0;
    }

    constructor(address contr) public {
        ownerAddress = msg.sender;
        _isContract(contr);
        delegateContract = Validators(contr);
    }

    function owner() public view returns (address){
        return ownerAddress;
    }

    modifier _isOwner {
        require(msg.sender == ownerAddress);
        _;
    }

    function changeDelegateContract(address contr) payable _isOwner public {
        delegateContract = Validators(contr);
    }

    function getValidatorList() public view returns (address[] memory) {
        return delegateContract.getValidatorsList();
    }

    function addValidator(address validator) public payable _isOwner {
        delegateContract.addValidator(validator);
    }
}
