# ldap filter expression
> Refer to: https://confluence.atlassian.com/kb/how-to-write-ldap-search-filters-792496933.html

## Matching all
以下语法都可以

    "(&(objectclass=*))"
    "(objectclass=*)"

## How do I match more than one attribute?
Notice the ampersand symbol '&' symbol at the start. Translated this means: search for objectClass=person AND objectClass=user.

    (&(objectClass=person)(objectClass=user))

Alternatively,The pipe symbol '|' denotes 'OR'.  Translated this means: search for objectClass=person OR object=user.

    (|(objectClass=person)(objectClass=user))

## Wildcards
    (&(objectClass=user)(cn=*Marketing*))

## Matching Components of Distinguished Names 

You may want to match part of a DN, for instance when you need to look for your groups in two subtrees of your server.

    // this will find groups with an OU component of their DN which is either 'Chicago' or 'Miami'.
    (&(objectClass=group)(|(ou:dn:=Chicago)(ou:dn:=Miami)))

## Using 'not'
To exclude entities which match an expression, use '!'.
So this will find all Chicago groups except those with a Wrigleyville OU component.

    (&(objectClass=group)(&(ou:dn:=Chicago)(!(ou:dn:=Wrigleyville))))

