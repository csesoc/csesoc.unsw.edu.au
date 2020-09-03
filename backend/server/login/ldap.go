/*
  LDAP
  --
  This part of the login module is responsible for checking whether a pair of
  credentials is valid by binding to Active Directory over LDAP, and then using
  the bind to search Active Directory for the name associated with the zID.
*/

package login

import (
	"fmt"
	"regexp"
	"github.com/go-ldap/ldap/v3"
)

// Define some constants we use for binding and searching.
const (
	ldapServer = "ad.unsw.edu.au:389"
	suffix = "@ad.unsw.edu.au"
	baseDN = "OU=IDM,DC=ad,DC=unsw,DC=edu,DC=au"
)

// bind opens a connection to the LDAP server and binds to it, returning
// a pointer to the connection, the authenticated state and any errors.
func bind(zId string, password string) (*ldap.Conn, bool, error) {
	// Open a TCP connection to the server.
	conn, err := ldap.Dial("tcp", ldapServer)
  
	if err != nil {
		return nil, false, fmt.Errorf("Could not connect. %s", err)
	}

	err = conn.Bind(zId + suffix, password)
  
	if err != nil {
	  	return conn, false, fmt.Errorf("Failed to bind. %s", err)
	}
	
	return conn, true, nil
}

// login authenticates by binding with credentials, returning the authentication
// state, and the name of the authenticated user.
func login(zId string, password string) (bool, string, error) {

	// We can expect a zId to be of the form zXXXXXXX.
	
	if valid, _:= regexp.MatchString(`[zZ]\d{7}`, zId); !valid {
		return false, "", fmt.Errorf("Invalid zID.")
	}

	conn, auth, err := bind(zId, password)

	// Ensure the connection is closed on completion.
	defer conn.Close()

	if !auth {
		// Return any failure to bind.
		return false, "", err
	}
	
	// Otherwise, the authentication was successful, so the bind can be used
	// to search for the user's name.

	result, err := conn.Search(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		10, // Max time for request (seconds).
		0, // Max size of request.
		false, // typesOnly flag 
		"(cn="+zId+")", // Common name is set to the zID, and is used as search filter.
		[]string{"displayName"}, // We only require the name of the user.
		nil, 
	))
	
	if err != nil || len(result.Entries) == 0{
		return true, "", err
	}

	// We assume zIDs are unique - that is, we only expect one entry for a 
	// given search.

	name := result.Entries[0].GetAttributeValue("displayName")

	return true, name, nil
}
