import { KeyInfo } from '@gno/api/gnomobiletypes_pb';

export type GnoConfig = {
  Remote: string;
  ChainID: string;
  KeyName: string;
  Password: string;
  GasFee: string;
  GasWanted: number;
  Mnemonic: string;
};

export type GnoAccount = KeyInfo;
