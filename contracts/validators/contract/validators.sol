pragma experimental ABIEncoderV2;

contract validators {
    address [] private validatorsAddressesArray =
    [0x64d4edefe8ba86d3588b213b0a053e7b910cad68,
    0x6a39b688d591ea00c9ea69658438794204b5cc62,
    0x4a14d788d86d021670ebcece1196631d66595984];

    function addValidator(address validatorAddress) public {
        validatorsAddressesArray.push(validatorAddress);
    }

    function getValidatorsList() public view returns (address[]){
        return validatorsAddressesArray;
    }
}
