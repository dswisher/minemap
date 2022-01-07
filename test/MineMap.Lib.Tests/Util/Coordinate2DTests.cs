// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.Collections.Generic;

using FluentAssertions;
using MineMap.Lib.Exceptions;
using MineMap.Lib.Util;
using Xunit;

namespace MineMap.Lib.Tests.Util
{
    public class Coordinate2DTests
    {
        public static IEnumerable<object[]> BadConversions()
        {
            yield return new object[] { (Action)(() => new Coordinate2D(CoordinateType2D.Region).ToChunk()) };
            yield return new object[] { (Action)(() => new Coordinate2D(CoordinateType2D.Region).ToChunkWithinRegion()) };
        }


        public static IEnumerable<object[]> IdentityConversions()
        {
            yield return new object[] { CoordinateType2D.Chunk, (Func<Coordinate2D, Coordinate2D>)(x => x.ToChunk()) };
            yield return new object[] { CoordinateType2D.ChunkWithinRegion, (Func<Coordinate2D, Coordinate2D>)(x => x.ToChunkWithinRegion()) };
            yield return new object[] { CoordinateType2D.Region, (Func<Coordinate2D, Coordinate2D>)(x => x.ToRegion()) };
        }


        [Fact]
        public void CanConstructDefault()
        {
            // Act
            var coord = new Coordinate2D();

            // Assert
            coord.X.Should().Be(0);
            coord.Z.Should().Be(0);
            coord.Type.Should().Be(CoordinateType2D.Block);
        }


        [Theory]
        [InlineData(CoordinateType2D.Chunk)]
        [InlineData(CoordinateType2D.Region)]
        public void CanConstructWithType(CoordinateType2D type)
        {
            // Act
            var coord = new Coordinate2D(type);

            // Assert
            coord.X.Should().Be(0);
            coord.Z.Should().Be(0);
            coord.Type.Should().Be(type);
        }


        [Theory]
        [InlineData(1, 2)]
        [InlineData(-3, 12)]
        public void CanConstructWithPosition(int x, int z)
        {
            // Act
            var coord = new Coordinate2D(x, z, CoordinateType2D.Block);

            // Assert
            coord.X.Should().Be(x);
            coord.Z.Should().Be(z);
        }


        [Fact]
        public void CanConvertToString()
        {
            // Arrange
            var coord = new Coordinate2D(1, 2, CoordinateType2D.Block);

            // Act
            var result = coord.ToString();

            // Assert
            result.Should().Contain("1");
            result.Should().Contain("2");
            result.Should().Contain("B");
        }


        [Theory]
        [InlineData(27, -15, 1, -1)]
        [InlineData(4, 8, 0, 0)]
        [InlineData(-4, 8, -1, 0)]
        [InlineData(4, -8, 0, -1)]
        [InlineData(-4, -8, -1, -1)]
        public void CanConvertBlockToChunk(int bx, int bz, int cx, int cz)
        {
            // Arrange
            var block = new Coordinate2D(bx, bz, CoordinateType2D.Block);

            // Act
            var chunk = block.ToChunk();

            // Assert
            chunk.X.Should().Be(cx);
            chunk.Z.Should().Be(cz);
            chunk.Type.Should().Be(CoordinateType2D.Chunk);
        }


        // TODO - add more test cases
        [Theory]
        [InlineData(4, 8, 0, 0)]
        public void CanConvertChunkToRegion(int cx, int cz, int rx, int rz)
        {
            // Arrange
            var chunk = new Coordinate2D(cx, cz, CoordinateType2D.Chunk);

            // Act
            var region = chunk.ToRegion();

            // Assert
            region.X.Should().Be(rx);
            region.Z.Should().Be(rz);
            region.Type.Should().Be(CoordinateType2D.Region);
        }


        // TODO - add more test cases
        [Theory]
        [InlineData(4, 8, 0, 0)]
        public void CanConvertBlockToRegion(int bx, int bz, int rx, int rz)
        {
            // Arrange
            var block = new Coordinate2D(bx, bz, CoordinateType2D.Block);

            // Act
            var region = block.ToRegion();

            // Assert
            region.X.Should().Be(rx);
            region.Z.Should().Be(rz);
            region.Type.Should().Be(CoordinateType2D.Region);
        }


        [Theory]
        [InlineData(4, 8, 4, 8)]
        [InlineData(33, 34, 1, 2)]
        [InlineData(-1, -2, 31, 30)]
        public void CanConvertChunkToChunkWithinRegion(int cx, int cz, int cix, int ciz)
        {
            // Arrange
            var block = new Coordinate2D(cx, cz, CoordinateType2D.Chunk);

            // Act
            var region = block.ToChunkWithinRegion();

            // Assert
            region.X.Should().Be(cix);
            region.Z.Should().Be(ciz);
            region.Type.Should().Be(CoordinateType2D.ChunkWithinRegion);
        }


        [Theory]
        [MemberData(nameof(IdentityConversions))]
        public void IdentityConversionsAreEasy(CoordinateType2D type, Func<Coordinate2D, Coordinate2D> act)
        {
            // Arrange
            var input = new Coordinate2D(type);

            // Act
            var output = act(input);

            // Assert
            output.Should().Be(input);
        }


        [Theory]
        [MemberData(nameof(BadConversions))]
        public void SomeConversionsAreNotPossible(Action act)
        {
            // Act and assert
            act.Should().Throw<CoordinateConversionException>();
        }
    }
}
