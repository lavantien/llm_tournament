import pygame
import sys
import numpy as np

G = 6.67408e-11
TIMESTEP = 60*60*24  # one day

class Body:
    def __init__(self, position, velocity, mass):
        self.position = np.array(position)
        self.velocity = np.array(velocity)
        self.mass = mass

def calculate_force(body1, body2):
    distance_vector = body1.position - body2.position
    distance = np.linalg.norm(distance_vector)
    force = G * body1.mass * body2.mass / distance**2
    force_vector = -force * distance_vector / distance
    return force_vector

def calculate_total_force(bodies, body):
    total_force = np.zeros(2)
    for b in bodies:
        if b != body:
            total_force += calculate_force(body, b)
    return total_force

def update_position_velocity(bodies, time_step):
    for body in bodies:
        force = calculate_total_force(bodies, body)
        acceleration = force / body.mass
        body.velocity += acceleration * time_step
        body.position += body.velocity * time_step

def draw(bodies, screen):
    screen.fill((0, 0, 0))
    for body in bodies:
        pygame.draw.circle(screen, (255, 255, 255), body.position.astype(int), int(body.mass**0.5))

def game_loop(bodies):
    pygame.init()
    size = width, height = 640, 480
    screen = pygame.display.set_mode(size)

    while True:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                sys.exit()

        update_position_velocity(bodies, TIMESTEP)
        draw(bodies, screen)
        pygame.display.flip()

bodies = [
    Body([100, 200], [0, 0], 1e30),  # sun
    Body([300, 300], [100, -100], 1e20),  # earth
    Body([350, 350], [120, -90], 1e20),  # moon
]

game_loop(bodies)

