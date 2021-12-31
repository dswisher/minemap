// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

namespace MineMap.Cli.Exceptions
{
    public class CliException : Exception
    {
        public CliException(string message)
            : base(message)
        {
        }
    }
}
