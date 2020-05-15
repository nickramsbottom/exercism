const colorCode = color => COLORS.indexOf(color);

const COLORS = ["black","brown","red","orange","yellow","green","blue","violet","grey","white"];

export const decodedValue = ([tens, ones]) => 10*colorCode(tens) + colorCode(ones);
