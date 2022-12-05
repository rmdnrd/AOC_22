"""AoC 2022 Day 5 Solution."""

def get_input(filename):
    """Read in file data."""
    with open(filename,'r',encoding='utf-8') as file:
        contents = file.read().splitlines()

    return contents

def get_moves(lines,index):
    """Separate move instructions."""
    moves = []
    for i in range(index,len(lines)):
        moves.append(lines[i])
    return moves

def get_starting_cargo(lines):
    """Determine starting cargo load."""
    cargo_start_index = 0
    move_index = 0
    cargo = {}

    # Find start of cargo moves
    i = 0
    while i < len(lines):
        if lines[i] == '':
            cargo_start_index = i-1
            move_index = i+1
            break
        i += 1

    # Load initial stacks
    for i in range (cargo_start_index,-1,-1):
        line = lines[i]

        # First line contains all cargo stacks
        if i == cargo_start_index:
            stacks = line.split()
            for stack in stacks:
                cargo.setdefault(int(stack),[])
            continue

        # use keys as stack index in cargo
        keys = list(cargo.keys())

        # Add Cargo
        for index in range(0,int(len(line)/4)+1):
            i = (index*4)+1
            temp = line[i:i+1]
            if temp != " ":
                cargo[keys[index]].append(temp)

    move_instructions = get_moves(lines,move_index)
    return cargo,move_instructions

def move_cargo_1(stacks,move_instructions):
    """Move cargo according to instructions."""
    for move in move_instructions:
        move_ct, move_fm, move_to = move.replace('move ', '').replace('from ', '').replace('to ', '').split(' ')
        move_ct = int(move_ct)
        move_fm = int(move_fm)
        move_to = int(move_to)
        for i in range(0,int(move_ct)):
            stacks[move_to].append(stacks[move_fm].pop(len(stacks[move_fm])-1))

    #return stacks
    solution_string = ""

    for cargo_list in stacks.items():
        solution_string += cargo_list[1][len(cargo_list[1])-1]

    #print(solution_string)
    return solution_string

def move_cargo_2(stacks,move_instructions):
    """A."""
    for move in move_instructions:
        move_ct, move_fm, move_to = move.replace('move ', '').replace('from ', '').replace('to ', '').split(' ')
        move_ct = int(move_ct)
        move_fm = int(move_fm)
        move_to = int(move_to)

        move_index = len(stacks[move_fm]) - move_ct
        for i in range(move_index,len(stacks[move_fm])):
            stacks[move_to].append(stacks[move_fm].pop(move_index))

    solution_string = ""
    for cargo_list in stacks.items():
        solution_string += cargo_list[1][len(cargo_list[1])-1]

    return solution_string

if __name__ == "__main__":

    text_lines = get_input("input-5.txt")

    cargo1,moves_list = get_starting_cargo(text_lines)
    answer1 = move_cargo_1(cargo1,moves_list)
    print(answer1)

    cargo2,moves_list = get_starting_cargo(text_lines)
    answer2 = move_cargo_2(cargo2,moves_list)
    print(answer2)
