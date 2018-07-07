int n = 0

proctype P(){
    do
    :: n < 10 ->
        n++
    ::else -> skip
    od
}
init{
    run P()
    run P()
    assert(n>2)
}