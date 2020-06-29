pragma solidity ^0.5.10;

contract parameters {
    string title;
    string version_hash;
    uint decimal = 10000;

    uint qTokenHoldersTransactionFeeShare;
    uint qTokenHoldersNativeAppsFeeShare;
    uint qTokenHoldersCoinbaseSubsidyShare;
    uint qidHoldersTransactionFeeShare;
    uint qidHoldersNativeAppsFeeShare;
    uint qidHoldersCoinbaseSubsidyShare;
    uint validatorNodesTransactionFeeShare;
    uint validatorNodesNativeAppsFeeShare;
    uint validatorNodesCoinbaseSubsidyShare;
    uint rootNodesTransactionFeeShare;
    uint rootNodesNativeAppsFeeShare;
    uint rootNodesCoinbaseSubsidyShare;

    constructor() public{
        title = "Q Constitution Parameter List";
        version_hash = "0BAB4BDC7C50E87F51273A1CD95D9588508349797C38432F788970693F40BF87";

        qTokenHoldersTransactionFeeShare = 3333;
        qTokenHoldersNativeAppsFeeShare = 3333;
        qTokenHoldersCoinbaseSubsidyShare = 3333;
        qidHoldersTransactionFeeShare = 3333;
        qidHoldersNativeAppsFeeShare = 3333;
        qidHoldersCoinbaseSubsidyShare = 3333;
        validatorNodesTransactionFeeShare = 3333;
        validatorNodesNativeAppsFeeShare = 3333;
        validatorNodesCoinbaseSubsidyShare = 3333;
        rootNodesTransactionFeeShare = 3333;
        rootNodesNativeAppsFeeShare = 3333;
        rootNodesCoinbaseSubsidyShare = 3333;
    }

    function setQTokenHoldersTransactionFeeShare(uint feeShare) public {
        qTokenHoldersTransactionFeeShare = feeShare;
    }

    function setQTokenHoldersNativeAppsFeeShare(uint feeShare) public {
        qTokenHoldersNativeAppsFeeShare = feeShare;
    }

    function setQTokenHoldersCoinbaseSubsidyShare(uint feeShare) public {
        qTokenHoldersCoinbaseSubsidyShare = feeShare;
    }

    function setQIDHoldersNativeAppsFeeShare(uint feeShare) public {
        qidHoldersNativeAppsFeeShare = feeShare;
    }

    function setQIDHoldersCoinbaseSubsidyShare(uint feeShare) public {
        qidHoldersCoinbaseSubsidyShare = feeShare;
    }

    function setQIDHoldersTransactionFeeShare(uint feeShare) public {
        qidHoldersTransactionFeeShare = feeShare;
    }

    function setValidatorNodesTransactionFeeShare(uint feeShare) public {
        validatorNodesTransactionFeeShare = feeShare;
    }

    function setValidatorNodesNativeAppsFeeShare(uint feeShare) public {
        validatorNodesNativeAppsFeeShare = feeShare;
    }

    function setValidatorNodesCoinbaseSubsidyShare(uint feeShare) public {
        validatorNodesCoinbaseSubsidyShare = feeShare;
    }

    function setRootNodesTransactionFeeShare(uint feeShare) public {
        rootNodesTransactionFeeShare = feeShare;
    }

    function setRootNodesNativeAppsFeeShare(uint feeShare) public {
        rootNodesNativeAppsFeeShare = feeShare;
    }

    function setRootNodesCoinbaseSubsidyShare(uint feeShare) public {
        rootNodesCoinbaseSubsidyShare = feeShare;
    }

    function getQTokenHoldersTransactionFeeShare() public view returns (uint) {
        return qTokenHoldersTransactionFeeShare;
    }

    function getQTokenHoldersNativeAppsFeeShare() public view returns (uint) {
        return qTokenHoldersNativeAppsFeeShare;
    }

    function getQTokenHoldersCoinbaseSubsidyShare() public view returns (uint) {
        return qTokenHoldersCoinbaseSubsidyShare;
    }

    function getQIDHoldersNativeAppsFeeShare() public view returns (uint) {
        return qidHoldersNativeAppsFeeShare;
    }

    function getQIDHoldersCoinbaseSubsidyShare() public view returns (uint) {
        return qidHoldersCoinbaseSubsidyShare;
    }

    function getQIDHoldersTransactionFeeShare() public view returns (uint) {
        return qidHoldersTransactionFeeShare;
    }

    function getValidatorNodesTransactionFeeShare() public view returns (uint) {
        return validatorNodesTransactionFeeShare;
    }

    function getValidatorNodesNativeAppsFeeShare() public view returns (uint) {
        return validatorNodesNativeAppsFeeShare;
    }

    function getValidatorNodesCoinbaseSubsidyShare() public view returns (uint){
        return validatorNodesCoinbaseSubsidyShare;
    }

    function getRootNodesTransactionFeeShare() public view returns (uint){
        return rootNodesTransactionFeeShare;
    }

    function getRootNodesNativeAppsFeeShare() public view returns (uint){
        return rootNodesNativeAppsFeeShare;
    }

    function getRootNodesCoinbaseSubsidyShare() public view returns (uint){
        return rootNodesCoinbaseSubsidyShare;
    }

    function getVersionHash() public view returns (string memory){
        return version_hash;
    }

    function setVersionHash(string memory hash) public {
        version_hash = hash;
    }

    function getDecimal()public view returns(uint){
        return decimal;
    }
}
