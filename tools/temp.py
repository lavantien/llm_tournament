import pygame

pygame.init()
width = 300
height = 600
window = pygame.display.set_mode((width, height))
pygame.display.set_caption('Tetris')

# Define grid dimensions
grid_size = 20
grid_width = width // grid_size
grid_height = height // grid_size
grid = [[0 for _ in range(grid_width)] for _ in range(grid_height)]

# Define colors for different tetrominos
colors = [
    (0, 0, 0),  # Empty cell
    (255, 0, 0),
    (0, 255, 0),
    (0, 0, 255),
    (255, 255, 0),
    (255, 127, 0),
    (0, 127, 255)
]

tetrominos = [
    [[1, 1, 1], [0, 1, 0]],
    [[0, 2, 2], [2, 2, 0]],
    [[3, 3, 3], [0, 3, 0]],
    [[4, 4, 0], [4, 4, 4]],
    [[0, 5, 5], [5, 5, 0]],
    [[6, 6, 6], [6, 6, 0]],
    [[7, 7, 7], [7, 0, 0]]
]


def rotate(matrix):
    return [list(x) for x in zip(*reversed(matrix))]


# Initialize the first piece
piece_type = 0
piece = tetrominos[piece_type]
piece_x = grid_width // 2 - len(piece[0]) // 2
piece_y = 0

score = 0
game_over = False

while not game_over:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            game_over = True
        elif event.type == pygame.KEYDOWN:
            if event.key == pygame.K_LEFT:
                piece_x -= 1
            elif event.key == pygame.K_RIGHT:
                piece_x += 1
            elif event.key == pygame.K_DOWN:
                piece_y += 1
            elif event.key == pygame.K_UP:
                piece = rotate(piece)

    # Check if the piece is out of bounds
    if piece_x < 0:
        piece_x = 0
    elif piece_x + len(piece[0]) > grid_width:
        piece_x = grid_width - len(piece[0])

    # Check if the piece collides with the bottom or other pieces
    for i, row in enumerate(piece):
        for j, cell in enumerate(row):
            if cell:
                if piece_y + i >= grid_height or grid[piece_y + i][piece_x + j]:
                    # The piece has landed, so add it to the grid and create a new piece
                    for i2, row2 in enumerate(piece):
                        for j2, cell2 in enumerate(row2):
                            if cell2:
                                grid[piece_y + i2][piece_x +
                                                   j2] = piece_type + 1
                    piece_type = (piece_type + 1) % len(tetrominos)
                    piece = tetrominos[piece_type]
                    piece_x = grid_width // 2
                    piece_y = 0
                    break

    # Move the piece down if it hasn't landed
    piece_y += 1
    if piece_y >= grid_height:
        piece_y = grid_height - 1
    for i, row in enumerate(grid):
        for j, cell in enumerate(grid[i]):
            if cell:
                if piece_y == i:
                    game_over = True

    window.fill((0, 0, 0))

    # Draw the game grid
    for i in range(grid_height):
        for j in range(grid_width):
            pygame.draw.rect(window, colors[grid[i][j]], pygame.Rect(
                j * grid_size, i * grid_size, grid_size - 1, grid_size - 1), 1)
            pygame.draw.rect(window, (0, 0, 255), pygame.Rect(
                j * grid_size, i * grid_size, grid_size, grid_size), 1)

    # Draw the current piece
    for i, row in enumerate.piece:
        for j, cell in row:
            if cell:
                pygame.draw.rect(window, colors[piece_type], pygame.Rect((piece_x + j) * grid_size, (piece_y + i) *
                                                                         grid_size, grid_size-1, grid_size-1))
                pygame.draw.rect(window, (255, 255, 0), pygame.Rect(
                    (piece_x + i) * grid_size, (piece_y) * grid_size, grid_size,  grid_size))

    font = pygame.font.SysFont('Arial', 40)
    score_text = font.render(f'Score: {score}', True, (255, 0, 0))
    window.blit(score_text, (width - score_text.get_width() - 10, 10))

    pygame.display.flip()

    clock = pygame.time.Clock()
    clock.tick(60)

pygame.quit()
