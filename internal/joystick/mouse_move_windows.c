// mouse_move_windows.c
#include <windows.h>
#include <math.h>

static float prevSpeedX = 0.0f;
static float prevSpeedY = 0.0f;

static void move_mouse_analog(int x, int y) {
    const int deadzone = 2000;
    const float maxSpeed = 10.0f;
    const float minSpeed = 1.2f;
    const float acceleration = 1.5f;
    const float deceleration = 0.0f;

    float deltaX = (float)x;
    float deltaY = (float)y;

    if (fabsf(deltaX) < deadzone) deltaX = 0.0f;
    if (fabsf(deltaY) < deadzone) deltaY = 0.0f;

    if (deltaX == 0.0f && deltaY == 0.0f) {
        prevSpeedX *= deceleration;
        prevSpeedY *= deceleration;

        if (fabsf(prevSpeedX) < 0.1f) prevSpeedX = 0.0f;
        if (fabsf(prevSpeedY) < 0.1f) prevSpeedY = 0.0f;

        POINT p;
        GetCursorPos(&p);
        SetCursorPos(p.x + (int)prevSpeedX, p.y + (int)prevSpeedY);
        return;
    }

    float length = hypotf(deltaX, deltaY);
    float normX = deltaX / length;
    float normY = deltaY / length;

    float intensity = fminf(length / 32767.0f, 1.0f);
    float targetSpeed = minSpeed + intensity * (maxSpeed - minSpeed);

    float newSpeedX = prevSpeedX + (normX * targetSpeed - prevSpeedX) * acceleration;
    float newSpeedY = prevSpeedY + (normY * targetSpeed - prevSpeedY) * acceleration;

    float speedLength = hypotf(newSpeedX, newSpeedY);
    if (speedLength > 0.0f) {
        float scale = targetSpeed / speedLength;
        newSpeedX *= scale;
        newSpeedY *= scale;
    }

    prevSpeedX = newSpeedX;
    prevSpeedY = newSpeedY;

    POINT p;
    GetCursorPos(&p);
    SetCursorPos(p.x + (int)newSpeedX, p.y + (int)newSpeedY);
}
