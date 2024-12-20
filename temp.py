import pygame
import math
import numpy as np

# Define colors
WHITE = (255, 255, 255)
BLUE = (0, 0, 255)

# Initialize pygame
pygame.init()

# Set the width and height of the screen [width,height]
size = (700, 500)
screen = pygame.display.set_mode(size)

pygame.display.set_caption("3-Body Problem Simulation")

# Loop until the user clicks the close button.
done = False

# Used to manage how fast the screen updates
clock = pygame.time.Clock()

# define the mass, size, and initial position & velocity of the three bodies
bodies = [
    {'mass': 10, 'size': 20, 'pos': [100.0, 100.0], 'vel': [1.0, 1.0]},
    {'mass': 20, 'size': 30, 'pos': [300.0, 300.0], 'vel': [0.0, 1.0]},
    {'mass': 30, 'size': 40, 'pos': [500.0, 200.0], 'vel': [0.5, 0.0]}
]

G = 10  # Gravitational Constant
dt = 0.1  # time step


# -------- Main Program Loop -----------
while not done:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            done = True

    screen.fill(WHITE)

    for body in bodies:
        pygame.draw.circle(screen, BLUE, [int(body['pos'][0]), int(body['pos'][1])], body['size'])

    for i in range(len(bodies)):
        for j in range(i+1, len(bodies)):
            r = np.linalg.norm(np.subtract(bodies[i]['pos'], bodies[j]['pos']))  # distance between i & j
            force_mag = G * bodies[i]['mass'] * bodies[j]['mass'] / (r*r)

            force_x = force_mag * (bodies[j]['pos'][0] - bodies[i]['pos'][0]) / r
            force_y = force_mag * (bodies[j]['pos'][1] - bodies[i]['pos'][1]) / r

            bodies[i]['vel'][0] += force_x * dt / bodies[i]['mass']
            bodies[i]['vel'][1] += force_y * dt / bodies[i]['mass']

            bodies[j]['vel'][0] -= force_x * dt / bodies[j]['mass']
            bodies[j]['vel'][1] -= force_y * dt / bodies[j]['mass']

    for body in bodies:
        body['pos'][0] += body['vel'][0] * dt
        body['pos'][1] += body['vel'][1] * dt

    pygame.display.flip()

    clock.tick(60)

pygame.quit()

