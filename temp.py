import pygame
import math
import sys

# Constants
WIDTH, HEIGHT = 800, 600
G = 1  # Gravitational constant
TIME_STEP = 0.5  # Time step for simulation
WHITE = (255, 255, 255)
BLACK = (0, 0, 0)

# Boundary limits
X_MIN, X_MAX = 0, WIDTH
Y_MIN, Y_MAX = 0, HEIGHT

# Body class


class Body:
    def __init__(self, mass, x, y, vx, vy, color):
        self.mass = mass
        self.x = x
        self.y = y
        self.vx = vx
        self.vy = vy
        self.color = color

    def draw(self, screen):
        pygame.draw.circle(screen, self.color, (int(self.x), int(self.y)), 5)

    def update(self, bodies):
        ax = 0
        ay = 0
        for other in bodies:
            if self != other:
                dx = other.x - self.x
                dy = other.y - self.y
                dist = math.hypot(dx, dy)
                force = G * self.mass * other.mass / (dist ** 2)
                ax += force * dx / (dist * self.mass)
                ay += force * dy / (dist * self.mass)
        self.vx += ax * TIME_STEP
        self.vy += ay * TIME_STEP
        self.x += self.vx * TIME_STEP
        self.y += self.vy * TIME_STEP

        # Boundary conditions
        if self.x < X_MIN:
            self.x = X_MIN
            self.vx = -self.vx * 0.8  # Dampen velocity a bit
        elif self.x > X_MAX:
            self.x = X_MAX
            self.vx = -self.vx * 0.8
        if self.y < Y_MIN:
            self.y = Y_MIN
            self.vy = -self.vy * 0.8
        elif self.y > Y_MAX:
            self.y = Y_MAX
            self.vy = -self.vy * 0.8


# Initialize Pygame
pygame.init()
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("3-Body Problem Simulation")

# Create bodies
bodies = [
    Body(1, 200, 300, 0, 1, (255, 0, 0)),  # Red body
    Body(1, 400, 300, 0, -1, (0, 255, 0)),  # Green body
    Body(1, 300, 200, 1, 0, (0, 0, 255))   # Blue body
]

# Main loop
running = True
while running:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    screen.fill(BLACK)

    for body in bodies:
        body.update(bodies)
        body.draw(screen)

    pygame.display.flip()
    pygame.time.delay(10)

pygame.quit()
sys.exit()
