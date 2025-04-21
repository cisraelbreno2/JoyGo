#include <windows.h>
#include <string.h>
#include <ctype.h>

static WORD getKeyCode(const char* keyName) {
    if (strlen(keyName) == 1) {
        char c = toupper(keyName[0]);
        return (WORD)c;
    }

    if (strcmp(keyName, "Enter") == 0) return VK_RETURN;
    if (strcmp(keyName, "Space") == 0) return VK_SPACE;
    if (strcmp(keyName, "Shift") == 0) return VK_SHIFT;
    if (strcmp(keyName, "Ctrl") == 0) return VK_CONTROL;
    if (strcmp(keyName, "Alt") == 0) return VK_MENU;
    if (strcmp(keyName, "Tab") == 0) return VK_TAB;
    if (strcmp(keyName, "Esc") == 0) return VK_ESCAPE;
    if (strcmp(keyName, "Backspace") == 0) return VK_BACK;
    if (strcmp(keyName, "Left") == 0) return VK_LEFT;
    if (strcmp(keyName, "Right") == 0) return VK_RIGHT;
    if (strcmp(keyName, "Up") == 0) return VK_UP;
    if (strcmp(keyName, "Down") == 0) return VK_DOWN;

    return 0;
}

static void keyDown(const char* keyName) {
    WORD keyCode = getKeyCode(keyName);
    if (keyCode != 0) {
        keybd_event((BYTE)keyCode, 0, 0, 0);
    }
}

static void keyUp(const char* keyName) {
    WORD keyCode = getKeyCode(keyName);
    if (keyCode != 0) {
        keybd_event((BYTE)keyCode, 0, KEYEVENTF_KEYUP, 0);
    }
}