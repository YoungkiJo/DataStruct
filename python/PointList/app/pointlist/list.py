from typing import List, Optional
from pydantic import BaseModel, Field

from pointlist.point import Point

class List(BaseModel):
    arr: List[int] = Field(default_factory=list)
    numOfData: int = 0
    curPosition: int = -1

    def LInsert(self, data: Point) -> None:
        self.arr.append(data)
        self.numOfData += 1

    def LFirst(self) -> bool:
        if self.numOfData == 0:
            return False

        self.curPosition = 0
        return True
    
    def LNext(self) -> bool:
        if self.curPosition >= self.numOfData -1:
            return False
        
        self.curPosition += 1
        return True
    
    def LRemove(self):
        rpos: int = self.curPosition
        num: int = self.numOfData

        for i in range(rpos, num-1):
            self.arr[i] = self.arr[i+1]
        
        self.numOfData -= 1
        self.curPosition -= 1
    
    def LCount(self) -> int:
        return self.numOfData