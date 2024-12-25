import pygame
import random
import math
import unittest

# Constants
WIDTH, HEIGHT = 800, 600
CENTER = (WIDTH // 2, HEIGHT // 2)
G = 1  # Gravitational constant (scaled for the simulation)
DT = 0.5  # Time step
FPS = 60
NUM_PAST_POSITIONS = 50  # Number of past positions to store for path tracing
COLLISION_DAMPING = 0.9  # Energy loss on collision

# Colors
WHITE = (255, 255, 255)
BLACK = (0, 0, 0)
RED = (255, 0, 0)
GREEN = (0, 255, 0)
BLUE = (0, 0, 255)


class Body:
    def __init__(self, x, y, mass, color, radius=10):
        self.x = x
        self.y = y
        self.mass = mass
        self.color = color
        self.radius = radius
        self.vx = 0
        self.vy = 0
        self.ax = 0
        self.ay = 0
        self.past_positions = []

    def calculate_force(self, other):
        dx = other.x - self.x
        dy = other.y - self.y
        distance = math.sqrt(dx**2 + dy**2)

        if distance == 0:
            return 0, 0

        force = G * self.mass * other.mass / (distance**2)
        fx = force * (dx / distance)
        fy = force * (dy / distance)
        return fx, fy

    def update_acceleration(self, bodies):
        self.ax = 0
        self.ay = 0
        for other in bodies:
            if other is not self:
                fx, fy = self.calculate_force(other)
                self.ax += fx / self.mass
                self.ay += fy / self.mass

    def update_velocity(self):
        self.vx += self.ax * DT
        self.vy += self.ay * DT

    def update_position(self):
        self.x += self.vx * DT + 0.5 * self.ax * DT**2
        self.y += self.vy * DT + 0.5 * self.ay * DT**2

        # Store past positions for path tracing
        self.past_positions.append((self.x, self.y))
        if len(self.past_positions) > NUM_PAST_POSITIONS:
            self.past_positions.pop(0)

    def handle_wall_collision(self):
        if self.x - self.radius < 0:
            self.x = self.radius
            self.vx = -self.vx * COLLISION_DAMPING
        elif self.x + self.radius > WIDTH:
            self.x = WIDTH - self.radius
            self.vx = -self.vx * COLLISION_DAMPING

        if self.y - self.radius < 0:
            self.y = self.radius
            self.vy = -self.vy * COLLISION_DAMPING
        elif self.y + self.radius > HEIGHT:
            self.y = HEIGHT - self.radius
            self.vy = -self.vy * COLLISION_DAMPING

    def handle_body_collision(self, other):
        dx = other.x - self.x
        dy = other.y - self.y
        distance = math.sqrt(dx**2 + dy**2)

        if distance < self.radius + other.radius:
            # Simple elastic collision response
            nx = dx / distance
            ny = dy / distance

            # Relative velocity
            dvx = other.vx - self.vx
            dvy = other.vy - self.vy

            # Dot product of relative velocity and normal vector
            vn = dvx * nx + dvy * ny

            # Bodies are moving towards each other
            if vn < 0:
                # Calculate impulse
                impulse = 2 * vn / (self.mass + other.mass)

                # Update velocities
                self.vx += impulse * other.mass * nx * COLLISION_DAMPING
                self.vy += impulse * other.mass * ny * COLLISION_DAMPING
                other.vx -= impulse * self.mass * nx * COLLISION_DAMPING
                other.vy -= impulse * self.mass * ny * COLLISION_DAMPING

                # Separate the bodies to prevent overlap
                overlap = (self.radius + other.radius - distance) / 2
                self.x -= overlap * nx
                self.y -= overlap * ny
                other.x += overlap * nx
                other.y += overlap * ny

    def draw(self, screen):
        pygame.draw.circle(screen, self.color,
                           (int(self.x), int(self.y)), self.radius)

    def draw_path(self, screen):
        if len(self.past_positions) > 1:
            pygame.draw.lines(screen, self.color, False,
                              self.past_positions, 2)


def create_bodies():
    bodies = []
    bodies.append(Body(CENTER[0] - 100, CENTER[1] -
                  50, random.uniform(50, 150), RED))
    bodies.append(Body(CENTER[0] + 100, CENTER[1] -
                  50, random.uniform(50, 150), GREEN))
    bodies.append(Body(CENTER[0], CENTER[1] + 100,
                  random.uniform(50, 150), BLUE))
    return bodies


def draw_walls(screen):
    pygame.draw.rect(screen, WHITE, (0, 0, WIDTH, HEIGHT), 2)


def main():
    pygame.init()
    screen = pygame.display.set_mode((WIDTH, HEIGHT))
    pygame.display.set_caption("3-Body Problem Simulation")
    clock = pygame.time.Clock()

    bodies = create_bodies()

    running = True
    while running:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False

        # Calculate forces and update accelerations
        for body in bodies:
            body.update_acceleration(bodies)

        # Update velocities and positions
        for body in bodies:
            body.update_velocity()
            body.update_position()

        # Handle collisions
        for i in range(len(bodies)):
            bodies[i].handle_wall_collision()
            for j in range(i + 1, len(bodies)):
                bodies[i].handle_body_collision(bodies[j])

        # Draw everything
        screen.fill(BLACK)
        draw_walls(screen)
        for body in bodies:
            body.draw_path(screen)
            body.draw(screen)

        pygame.display.flip()
        clock.tick(FPS)

    pygame.quit()


class TestBodyMethods(unittest.TestCase):
    def setUp(self):
        self.body1 = Body(100, 100, 50, RED)
        self.body2 = Body(200, 200, 100, BLUE)

    def test_calculate_force(self):
        fx, fy = self.body1.calculate_force(self.body2)
        self.assertNotEqual(fx, 0)
        self.assertNotEqual(fy, 0)

    def test_update_acceleration(self):
        bodies = [self.body1, self.body2]
        self.body1.update_acceleration(bodies)
        self.assertNotEqual(self.body1.ax, 0)
        self.assertNotEqual(self.body1.ay, 0)

    def test_update_velocity(self):
        self.body1.ax = 1
        self.body1.ay = 1
        self.body1.update_velocity()
        self.assertNotEqual(self.body1.vx, 0)
        self.assertNotEqual(self.body1.vy, 0)

    def test_update_position(self):
        self.body1.vx = 1
        self.body1.vy = 1
        self.body1.update_position()
        self.assertNotEqual(self.body1.x, 100)
        self.assertNotEqual(self.body1.y, 100)

    def test_handle_wall_collision(self):
        self.body1.x = self.body1.radius - 1
        self.body1.vx = -1
        self.body1.handle_wall_collision()
        self.assertEqual(self.body1.x, self.body1.radius)
        self.assertGreater(self.body1.vx, 0)

    def test_handle_body_collision(self):
        self.body1.x = 100
        self.body1.y = 100
        self.body2.x = 100 + self.body1.radius + self.body2.radius - 1
        self.body2.y = 100
        self.body1.vx = 1
        self.body2.vx = -1
        self.body1.handle_body_collision(self.body2)
        self.assertLess(self.body1.vx, 0)
        self.assertGreater(self.body2.vx, 0)


if __name__ == "__main__":
    # unittest.main() # Uncomment to run unit tests
    main()
