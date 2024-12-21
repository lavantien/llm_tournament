import pygame
import math

# Constants
G = 6.67430e-11  # Gravitational constant
M_EARTH = 5.972e24  # Mass of the Earth in kg
M_MARS = 6.42e23  # Mass of Mars in kg
M_JUPITER = 1.898e27  # Mass of Jupiter in kg

# Colors
WHITE = (255, 255, 255)
RED = (255, 0, 0)
GREEN = (0, 255, 0)
BLUE = (0, 0, 255)

class Body:
    def __init__(self, x, y, vx, vy, mass, color):
        self.x = x
        self.y = y
        self.vx = vx
        self.vy = vy
        self.mass = mass
        self.color = color

    def update(self):
        self.x += self.vx
        self.y += self.vy

    def draw(self, screen):
        pygame.draw.circle(screen, self.color, (int(self.x), int(self.y)), 5)

def calculate_force(body1, body2):
    dx = body2.x - body1.x
    dy = body2.y - body1.y
    distance = math.sqrt(dx**2 + dy**2)
    if distance == 0:
        return 0, 0
    force = G * body1.mass * body2.mass / distance**2
    angle = math.atan2(dy, dx)
    fx = force * math.cos(angle)
    fy = force * math.sin(angle)
    return fx, fy

def simulate(bodies):
    for i, body1 in enumerate(bodies):
        fx, fy = 0, 0
        for j, body2 in enumerate(bodies):
            if i != j:
                fx1, fy1 = calculate_force(body1, body2)
                fx += fx1 / body1.mass
                fy += fy1 / body1.mass
        body1.vx += fx
        body1.vy += fy

def main():
    pygame.init()
    screen_width, screen_height = 800, 600
    screen = pygame.display.set_mode((screen_width, screen_height))
    clock = pygame.time.Clock()

    bodies = [
        Body(100, 100, 0, 0, M_EARTH, RED),
        Body(300, 300, 0, 0, M_MARS, GREEN),
        Body(500, 500, 0, 0, M_JUPITER, BLUE)
    ]

    for i in range(1000):
        simulate(bodies)

    running = True
    while running:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False

        screen.fill(WHITE)

        for body in bodies:
            body.update()
            body.draw(screen)

        pygame.display.flip()
        clock.tick(60)

    pygame.quit()

if __name__ == "__main__":
    main()

