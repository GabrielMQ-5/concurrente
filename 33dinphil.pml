#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

#define N 5

byte fork[N] = { 1, 1, 1, 1, 1 }
byte room = 4

active[N] proctype P() {
    byte izq = _pid
    byte der = (_pid + 1) % N
    byte i
    for (i : 1..10) {
        printf("Filósofo %d está pensando\n", _pid)
        wait(room)
        wait(fork[izq])
        wait(fork[der])
        printf("Filósofo %d está comiendo\n", _pid)
        signal(fork[izq])
        signal(fork[der])
        signal(room)
    }
}