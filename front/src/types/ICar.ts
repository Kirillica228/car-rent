import IBrand from "./IBrand";
import IImages from "./IImage";
import IType from "./IType";

interface ICar {
  id: number;
  brand: IBrand;
  type: IType;
  model: string;
  year: number;
  licensePlate: string;
  status: string;
  price: number;
  isVisible: boolean;
  images: IImages[];
  color: string;
  description: string;
  transmission: string;
};

export default ICar;