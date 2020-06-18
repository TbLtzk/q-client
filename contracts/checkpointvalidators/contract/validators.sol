pragma experimental ABIEncoderV2;

contract validators {
    string[] private validatorsPubkeysArray =
    ["0a84e562c9b22e43269b7dca215cf2ed8c20bbf35da67bae8d5ee81b36d8bbb69e3ec704b9b6f7501059fe861843a836b2fbab641f36616cdd77365b1a522d5b"];

    function addValidator(string validatorPubkey) public {
        validatorsPubkeysArray.push(validatorPubkey);
    }

    function getValidators() public view returns (string[]){
        return validatorsPubkeysArray;
    }
}
