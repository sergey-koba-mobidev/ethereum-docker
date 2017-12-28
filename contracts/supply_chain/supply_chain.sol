contract SupplyChain {
    string public name;
    string public description;
    address owner;

    struct Product {
        uint uid;
        uint state;
        string data;
        address creator;
        address owner;
        bool initialized;
    }

    mapping (uint => Product) products;
    uint public numberOfProducts;

    function SupplyChain(string _name, string _description) public {
        name = _name;
        description = _description;
        numberOfProducts = 0;
    }

    function addProduct(string data) public returns (bool success) {
        numberOfProducts++;
        products[numberOfProducts] = Product(numberOfProducts, 0, data, msg.sender, msg.sender, true);
        return true;
    }

    function transferProduct(uint uid, address addr) public returns (bool success) {
        require(products[uid].initialized);
        require(products[uid].owner == msg.sender);
        products[uid].owner = addr;
        return true;
    }

    function getProductInfo(uint _uid) public constant returns (uint uid, uint state, string data, address owner, address creator) {
        var product = products[_uid];
        require(product.initialized);
        return (product.uid, product.state, product.data, product.owner, product.creator);
    }
}

// TODO: Administrator access
// TODO: split into many contracts to track transactions
// TODO: Product id should be returned via event