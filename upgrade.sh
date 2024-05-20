ag -lF 'golang.org/x/net v0.17.0' | ag 'go.mod$' |xargs -n1 gsed -i 's#golang.org/x/net v0.17.0#golang.org/x/net v0.25.0#g' 
