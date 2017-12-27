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
    }

    mapping (uint => Product) products;
    uint public numberOfProducts;

    function SupplyChain(string _name, string _description) public {
        name = _name;
        description = _description;
        numberOfProducts = 0;
    }

    function addProduct(string data) public returns (uint uid) {
        numberOfProducts++;
        products[numberOfProducts] = Product(numberOfProducts, 0, data, msg.sender, msg.sender);
        return numberOfProducts;
    }

    function transferProduct(uint uid, address addr) public returns (bool success) {
        if (!products[uid].uid) {
            throw;
        }
        if (products[uid].owner != msg.sender) {
            throw;
        }
        products[uid].owner = addr;
        return true;
    }

    function getProductInfo(uint uid) public constant returns (uint uid, uint state, string data, address owner, address creator) {
        var product = products[uid];
        if (!product.uid) {
            throw;
        }
        return (product.uid, product.state, product.data, product.owner, product.creator);
    }
}