#!/usr/bin/env python3
 
import time
import argparse

def main(): 
    parser = argparse.ArgumentParser(description="MB of memory to consume")
    parser.add_argument("-m", type=int, default=128, help="Amt of mem")
    args = parser.parse_args()
    mem_alloc = args.m

    memory_size = mem_alloc * 1024 * 1024

    print(f"Nom nom nom {mem_alloc:.0f} MB memory")
    memory = bytearray(memory_size)

    for i in range(0, memory_size, 1024):
        memory[i] = 1

    print("Ctrl-C to exit")
    try:
        while True:
            time.sleep(1)
    except:
        print("Exiting...")

if __name__ == "__main__":
    main()

