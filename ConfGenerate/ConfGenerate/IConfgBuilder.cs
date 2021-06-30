using System.Collections.Generic;

namespace ConfGenerate
{
    public interface IConfgBuilder
    {
        public virtual List<string> Read();
    }
}