import { useState } from "react";

const useSelectOne = (): [number | null, (id: number) => void, () => void] => {
  const [selected, setSelected] = useState<number | null>(null);
  const onSelect = (id: number) => setSelected(id);
  const unSelect = () => setSelected(null);

  return [selected, onSelect, unSelect];
};

export default useSelectOne;
