import { FunctionComponent, ReactNode } from "react";

type Props = {
  when: boolean;
  show: ReactNode;
  otherwise?: ReactNode;
};

const Render: FunctionComponent<Props> = ({ when, show, otherwise }) => {
  if (when) {
    return show;
  }
  return otherwise;
};

export default Render;
