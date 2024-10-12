from pydantic import BaseModel

class Point(BaseModel):
    xpos: int = 0
    ypos: int = 0

    def SetPointPos(self, xpos: int, ypos: int):
        self.xpos = xpos
        self.ypos = ypos

    def ShowPointPos(self):
        print(f"[{self.xpos}, {self.ypos}]")    

def PointComp(pos1: Point, pos2: Point):
    if pos1.xpos == pos2.xpos and pos1.ypos == pos2.ypos:
        return 0
    elif pos1.xpos == pos2.xpos:
        return 1
    elif pos1.ypos == pos2.ypos:
        return 2
    else:
        return -1