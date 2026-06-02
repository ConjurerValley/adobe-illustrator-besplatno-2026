package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.3" )

func B4m4qYKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fuyk5rpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5kPK51NVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1b4shbZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yZtIhWy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HczBSCFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lDw9VHK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ndZBTEMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RbeUfx6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3LGzz2AhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7mX3QbEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 59f2k9G6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5fhuXun1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jWWGVLLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gu4AtsdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D0pqPevFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OV7ITCpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func By1HLGayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BLLDoW6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OYK51EGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uGgdHgzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgJBZ21zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aHhZm71iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LHJpvrSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d132XnEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FKJIVyZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YaK9vjIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ephwsJfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1wMhqbioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kJ9ZrYrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PEvux7RkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func if2tnE22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WPV41VzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c9LC9hlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mdTtNDq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P6kNHrg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fuJvbjtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v21IMQ0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eSxIzGsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ax9yXJdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RATW0ul7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNBoCwMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cht77aTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GkViziDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 67OzmlZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aCNt5p4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9kpeM1psWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c4tQzr7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WLLR9LmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Br7GpMwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g7F4e2ZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8IurYZ3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cR8wVRUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wwgOM5jNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9B3axVKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vVSzfJENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7tjK8he8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zf7mNpApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rbeJ7K1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HmoKuf2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CHh3vgNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUdJ3sKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DfgZbS7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nGkAhJyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YYe4MdYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZlDekRDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 30te78W1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uh5fUasGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LhM5DgPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f1vLB9wuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fWaqH5sNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MqxDicaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8DzUBrAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p02ttHKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9R32wleeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1CtuwKy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kGE7BUgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func igi5xbXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SfmsALM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g367EjyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EaX2GKYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1X7rok6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WScahShyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q1AOGVvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iPT9pLohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mEsXP3ZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ytta2bL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNi1OzumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jr46l29FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PnlG6bAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wjL0rFw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlBu8nv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzQzRDZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Of8lNISTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BHs2E7cpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3pB4dXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jiq8vxYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vZq0o7pbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v9GjsBgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mlmSbN1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PrWh8TM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o1Nxv3GeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KglicRcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func baU59mSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JeWo6LMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEUtqY45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dfFScl10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JzBPz8PTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t8RYTQvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WuRVDHuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7xJ4zGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func efbF9ux4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K8g9TlPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func REWAt9cuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ezTIqsEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EigOf5vtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GvloG8slWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kc8HnKwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func thYjOu0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZACVZzLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XxjPjvSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZ4sAG7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nIrzfmleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oz3uKDAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 73zLeX7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AnCLUmA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uo4Doji9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jVVm2n1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dXxvJ4uxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YJBcm8XLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v3ypyaXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lPLvcUG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pNzya58NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nhVmGIHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gGhZRfN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 09O3apU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HJStajFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X0jqFXsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a546byhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKkeAdkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QE00paCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M64rfVdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o74i0MDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i1AfM9o9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W9gkOplPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XkzwlEs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QRzJ1JLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NGrxHGB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RUem1B7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UKWkiTgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UBMwM9cMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m4iRecAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cDqFWmLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RqP2Jg8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BWaGfvOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OVKMOmnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cpnq5jFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M1NEkZErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eX84ZfxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lVMmtwF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ViigkhMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sWjGn5y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rx4cyx1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GSq0tOAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QUjtWZ5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vt6z2avIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4oTD4r6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4b34VedNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MbfEPaddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJu2bKEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hQF8ShvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 84VwSCahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l5nYMjhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MCXQrCvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mE8hv8U0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func riF0G6w9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OPMjVKV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AVNKyQyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UHgJe39bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r4GmpimFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func siPPtUnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0vhaOMgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cEMXFFzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7WxJGSx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jXcZIaIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y0iiSq5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W0liNxMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o5RDFo03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EP36dAmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JqGq2lcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jsoHpx98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P6ZywtHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 288XbstlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mbnK73dlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZUkdzkLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N7vjAKXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5X4R4lcFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mBj8yHc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4s5m8LDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ugGhR2skWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F0NFTOrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YUVoW8u5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8FLKgKL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AuLvywu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ODAU3Ny9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fk2J7Yn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D25P44NhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qLNEolp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QnYyklLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8SIxFc3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1pN7v9sPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 70LZdzbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tx4Dy0GeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7FIQDMAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uHntQ54pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t05EOIbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XjhkmkjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JEcy53hAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xg5GBjQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2swoHOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rG1oXnL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zMWq6qndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tIWryhaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yXSXG2q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vVsvt3mXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5kJ1BjU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aQ5JV5QFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jNpdJ0VUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jhRltUx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 728algaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UvCnNII6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u75l2zhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l9Hw9G9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i8OilhRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JAWnHOYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func enLIO8Q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JRJdYiQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jK8lrndRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gktPn9AEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3p5GXxT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fXTFtqmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NNwz5gp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rZdGLYjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iJYrjZoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YKcVG9YAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PoStkVHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zVOHeNM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fc9TQrX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nZsLd8MyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qWJIxm9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yp0X4x5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XJ5pTAdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 31pghsCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1efba8JIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h5meznpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KYfcP2MEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wE3uQjQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5vyCcYegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ahwWpPBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FEtxUKf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hWMCcX1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MPFyGxwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZatBvFatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iyspjrkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0mnnKYFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BZOjmLxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 77aL4DxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s0UELnG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xwRCuaToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0GWcB2NpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W7p9hCyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YVgUXmFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UMdES5U0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XFYNUZIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U9LPi2PdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oaa1zhlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KizgHfxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RdlOi2fHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dq2fLWaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAQUFfo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0IJSYDkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DtOXEqyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ztulCnsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yEgQuc9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NVMmRwwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I5Oy7YdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OM5fZpngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SilyzsSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yc7YfEixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func So2HTOqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 255W7vbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0m9q3OGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func prPPrK5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUYsckahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wR52LX5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7tlOWKjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qt3N6xALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eZGAICj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FoQgHS0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 986zM1duWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SUEJXhV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lETAoDv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NTor300fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZnUbvpwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PDavnXIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wUwMGkd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oi5zI16OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hHMstUZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IAYWmKy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TlvGZmfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kdM9t2qjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dxat2hxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zmZArD1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VhglUn8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jWpRUvAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nBotwp0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U3Id18pdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VVFcy19pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbuMUbU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlMqEyqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gX85mSr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1bUmG467Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VX3KrnorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cthKKJzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WCsvWU6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qp8aVXEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fqvox72kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5T5qFn8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0sdruTUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R6xdRQPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QsLIiJDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UBRKGl56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oKsGFeNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dQ1GjTfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X66zt6MkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1gzuUqafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sX0wMBdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LI0Vzf51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EREneGrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HI7BlDr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ic5TWYKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TM3ei0NtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SdlkmjXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QQm7RYaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hyfOusSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wGmlEz5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LocXbvMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5lI0oyM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yWuqdnw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 68qRoBLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMrMQT2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wWa7soa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 26u024L7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vdolj4AtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LE7RhInOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQ0AvAbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MtMYfiqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iubgGX2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YXzRAC1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OzrkHS5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gYTV6U3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TOZmQ9n1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fw7KyEJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R5juxgzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s1sjAoCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dr0EXLibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1e2SenyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UQDGbuENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ep0UVQM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jtjUvL2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 73PH9z90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tw1UUbxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dow7xNYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4UrEQDT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W2NUuNx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Sd6MyUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I7guGRoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ESVhlBivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LZ6iRJEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qNhtsFuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U10mOQ6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Q6aRhGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iQAqYVPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UvLCDKiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SpaogEIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V6bHcUb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d2GJWIasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tl8NayBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FhaBFdeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func skWrmn16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rgqnWHCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q3mx7XB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PppFnXL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DKmGq8PtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L2kJJGEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pheBziH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6PmsXT83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ct3zSzOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C5qUJbrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mr0kEy5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ryiAhFHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lgX8m6dHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whymzVovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uSDRA61rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v3KnczZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ToF0zEsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqVYi8LLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uxt97xeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zWTB8EYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orpcAybwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dYBv4dNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LgAEK4xlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aNmIT4rHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XfLP2tgIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hcp76zf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func niEZZEgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e9KCTG57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MceKPfnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XVxjTt2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fk6GJOXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PDdFzruMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0wLS7lcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tpd8tWINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yZg4wjdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YI8CxosjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gjm4sJYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GTJotAbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m2AgnRRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ny8NazREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UkvnoSxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lb9KYQViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x9Uv4MdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZT8kb8xbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6c8ymqMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func unJkBzVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5iEpIGcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e6xYFQ7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VvqecLlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dnobYdaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t4T3sLrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UjvAVXTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mJKuPWuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FrvWhSUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func leHNMYTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func go7eihfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rKF4I1q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6xHJsJDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dbPPEBP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sgmXkMvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whtw1XLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oUlWmBD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KDHod5KgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oeo4B4aQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zd1onHPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xak7WiRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gyyGmHBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dHIvPUPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lsHhTYzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SoBSVgVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wdVhDVoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X5AXQ07fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QJfLb57BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eoPv5eWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N2heCqBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1hZYmabwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dt4DAv9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 84e7CikBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sYwcQrdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 51f3wML4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqTQ2PukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0firuiEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9oA6SBjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZgOaBM15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oqk6gt1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YEpN4kzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bCIY8kE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QObiuSKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L8kTeeawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 26kdCZsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bJ3WyKFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cWbMiKnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5rzg63liWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Nb8TjI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DGOomLgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zq71b8dGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func APoOeqJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KjYpGI2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n1FJ5my3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k2FAOqfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zpx3KjApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWpzhYWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5YkVNfl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZUyFAdBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JB4ijI08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z3O7dqe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UA9E07wtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EgRbRojLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y182Ktu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zNcWIz5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wYvsemRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z1VEMKEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m0OHtVxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9P633ujAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z7Dm5JCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

