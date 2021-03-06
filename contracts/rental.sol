// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;
/// @title property listing and rental contract 
contract rentalContract {
    
    struct property {

        address owner;
        uint id;
        uint perNight;
        bool available;
        uint securityDeposit;

        // array of all rentals 
        uint[] t1;
        uint[] t2;

    }

    struct rental {
        
        address owner;
        uint id;
        address renter;

        // check in check out time of renter
        uint t1;
        uint t2;
    
    }

    struct securityDep {

        address owner;
        uint id;
        address renter;

        // renter will be able to withdraw deposit after certain amount of time 
        uint amount;

        // this will be used to release after x number of days 
        uint timestamp;

        // default is false
        bool dispute;

    }


    // property struct mapping 
    mapping(address => mapping (uint => property)) public properties;

    // rental struct mapping 
    mapping(address => mapping (uint => rental)) public rentals;

    // allows for multiple property listings per address @dev potentially can be optimized
    mapping(address => uint[]) private listmappingOwner;
    mapping(address => uint[]) private listmappingRenter;

    // security deposit mapping 
    mapping(address => mapping (uint => securityDep)) public securityDeposits;

    // events are not yet implemented in the Solidity Go binding 
    event PropertyListed(address owner, uint ID);
    event PropertyRented(address owner, uint ID, address renter, uint checkin, uint checkout);
    event PriceUpdated(address owner, uint newPrice, uint ID);

    // used by listProperty and rentProperty
    uint private ID;
    
    function listProperty(uint perNight, uint securityDeposit) public {
        
        ID = listmappingOwner[msg.sender].length;
        
        properties[msg.sender][ID].owner = msg.sender;
        properties[msg.sender][ID].id = ID;
        properties[msg.sender][ID].perNight = perNight;
        properties[msg.sender][ID].available = true;

        properties[msg.sender][ID].securityDeposit = securityDeposit;
        listmappingOwner[msg.sender].push(ID);
        emit PropertyListed(msg.sender, ID);

    }

    // handle multiple rentals per renter 
    mapping (address => uint[]) private previousRentals;

    function rentProperty(address payable owner, uint id, uint t1, uint t2) public payable {
        
        uint start;
        uint end;
        uint duration;
    
        uint fee;
        uint payment;
        uint securityDeposit;
        
    
        require(id <= (listmappingOwner[owner].length - 1), "ID not found");
        duration = (t2 - t1);
        
        // minimum rental duration is 1 day
        require(duration <= 86400);

        // for loop to check if rental dates overlap with other renters 
        for (uint i=0; i<properties[owner][id].t1.length; i++) {
            
            start = properties[owner][id].t1[i];
            end = properties[owner][id].t2[i];
        
            require((end <= t1 || start >= t2) == true);

        }

        securityDeposit = properties[owner][id].securityDeposit;

        fee = (duration * properties[owner][id].perNight) / 86400 + properties[owner][id].securityDeposit;

        payment = (duration * properties[owner][id].perNight) / 86400;

        require (msg.value >= fee);

        // rental and security deposits struct is mapped to renter address not owner address 

        // id is of owner properties 
        // ID is of renter rentals 

        // push to rentals struct 
        ID = listmappingRenter[msg.sender].length;

        rentals[msg.sender][ID].owner = owner;
        rentals[msg.sender][ID].id = id;
        rentals[msg.sender][ID].renter = msg.sender;
        rentals[msg.sender][ID].t1 = t1;
        rentals[msg.sender][ID].t2 = t2;

        // push check in and check out to properties struct 
        properties[owner][id].t1.push(t1);
        properties[owner][id].t2.push(t2);

        // security Deposit struct 
        securityDeposits[msg.sender][ID].owner = owner;
        securityDeposits[msg.sender][ID].id = id;
        securityDeposits[msg.sender][ID].renter = msg.sender;
        securityDeposits[msg.sender][ID].amount += securityDeposit;
        securityDeposits[msg.sender][ID].dispute = false;
        securityDeposits[msg.sender][ID].timestamp = block.timestamp;

        listmappingRenter[msg.sender].push(ID);

        owner.transfer(payment);
        
        emit PropertyRented(owner, id, msg.sender, t1, t2);
    
    }


    // update price per night 
    function updateRentalPrice(uint newPrice, uint id) public { 

        properties[msg.sender][id].perNight = newPrice;
        
        emit PriceUpdated(msg.sender, newPrice, id);

    }


    function UpdatePropertyAvailability(uint id, bool available) public {

        properties[msg.sender][id].available = available;

    }

    // in future members of DAO can vote to resolve dispute 
    function fileDispute(address owner, uint id, bool dispute) public {

        require(msg.sender == securityDeposits[owner][id].owner);

        securityDeposits[msg.sender][id].dispute = dispute;

    }
    
    // user is renter
    function releaseDeposit(address payable renter, uint id) public {

        uint amount;
        
        require(msg.sender == securityDeposits[renter][id].owner || msg.sender == securityDeposits[renter][id].renter);
        
        // release after a week if there is no dispute
        require(securityDeposits[renter][id].timestamp + 604800 < block.timestamp, "Deposit still on hold");
        require(securityDeposits[renter][id].dispute == false, "Owner has filed a dispute");
   
        amount = securityDeposits[renter][id].amount;
        renter.transfer(amount);
        
    }

}
