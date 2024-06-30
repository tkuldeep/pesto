import { useEffect, useState } from "react";
import { Task } from "../../App";
import { Button } from "../ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "../ui/dialog";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { X } from "lucide-react";

interface Props {
  data: Task | null;
  handleEditSave: (title: string, desc: string) => void;
  hanldeClose: () => void;
}

export const DialogComp = ({ data, handleEditSave, hanldeClose }: Props) => {
  const [inputState, setInputState] = useState({
    title: data?.title || "",
    desc: data?.desc || "",
  });

  useEffect(() => {
    if (data)
      setInputState({
        title: data.title,
        desc: data.desc,
      });
  }, [data]);

  return (
    Boolean(data) && (
      <Dialog open={true}>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader className="flex flex-row justify-between items-start">
            <div>
              <DialogTitle>Edit Task</DialogTitle>
              <DialogDescription>
                Make changes to your task here. Click save when you're done.
              </DialogDescription>
            </div>
            <X onClick={hanldeClose} className="cursor-pointer mt-0" />
          </DialogHeader>
          <div className="grid gap-4 py-4">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="title" className="text-right">
                Title
              </Label>
              <Input
                id="title"
                value={inputState.title}
                className="col-span-3"
                onChange={(e) => {
                  setInputState((prev) => ({ ...prev, title: e.target.value }));
                }}
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="description" className="text-right">
                Description
              </Label>
              <Input
                id="description"
                value={inputState.desc}
                className="col-span-3"
                onChange={(e) => {
                  setInputState((prev) => ({ ...prev, desc: e.target.value }));
                }}
              />
            </div>
          </div>
          <DialogFooter>
            <Button
              type="button"
              onClick={() => handleEditSave(inputState.title, inputState.desc)}
              disabled={!inputState.title.length || !inputState.desc.length}
            >
              Save changes
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    )
  );
};
